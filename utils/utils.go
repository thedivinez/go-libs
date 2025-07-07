package utils

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/big"
	"mime/multipart"
	"slices"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func Encrypt(key, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", errors.WithStack(err)
	}
	b := base64.StdEncoding.EncodeToString([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", errors.WithStack(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(key, text string) (string, error) {
	textBytes, err := hex.DecodeString(text)
	if err != nil {
		return "", errors.WithStack(err)
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", errors.WithStack(err)
	}
	if len(textBytes) < aes.BlockSize {
		return "", errors.New("failed to decrypt cipher")
	}
	iv := textBytes[:aes.BlockSize]
	textBytes = textBytes[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(textBytes, textBytes)
	data, err := base64.StdEncoding.DecodeString(string(textBytes))
	if err != nil {
		return "", errors.WithStack(err)
	}
	bs, err := hex.DecodeString(hex.EncodeToString(data))
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(bs), nil
}

func GenerateRandomNumber(numberOfDigits int) (int, error) {
	maxLimit := int64(int(math.Pow10(numberOfDigits)) - 1)
	lowLimit := int(math.Pow10(numberOfDigits - 1))
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxLimit))
	if err != nil {
		return 0, errors.WithStack(err)
	}
	randomNumberInt := int(randomNumber.Int64())
	// Handling integers between 0, 10^(n-1) .. for n=4, handling cases between (0, 999)
	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}
	// Never likely to occur, kust for safe side.
	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}
	return randomNumberInt, nil
}

func Transcode(in, out any) error {
	resultBytes, err := json.Marshal(in)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(json.Unmarshal(resultBytes, &out))
}

func Today() time.Time {
	now := time.Now().UTC()
	year, month, day := now.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func Yesterday() time.Time {
	now := time.Now().UTC()
	year, month, day := now.Date()
	return time.Date(year, month, day-1, 0, 0, 0, 0, time.UTC)
}

func Tomorrow() time.Time {
	now := time.Now().UTC()
	year, month, day := now.Date()
	return time.Date(year, month, day+1, 0, 0, 0, 0, time.UTC)
}

type Countdown struct {
	T float64 // total remaning time in seconds
	D int64   // days
	H int64   // hours
	M int64   // minutes
	S int64   // seconds
	P float64 // percentage of total time remaining
	F string  // formatted time
}

func getTimeString(remaining Countdown) string {
	res := fmt.Sprintf("%02d:%02d:%02d:%02d", remaining.D, remaining.H, remaining.M, remaining.S)
	if remaining.D <= 0 && remaining.H > 0 {
		res = fmt.Sprintf("%02d:%02d:%02d", remaining.H, remaining.M, remaining.S)
	} else if remaining.D <= 0 && remaining.H <= 0 {
		res = fmt.Sprintf("%02d:%02d", remaining.M, remaining.S)
	}
	return res
}

func getTimeRemaining(to time.Time) Countdown {
	total := time.Until(to).Seconds()
	countDown := Countdown{
		T: total,
		M: int64(total/60) % 60,
		S: int64(int(total) % 60),
		D: int64(total / (60 * 60 * 24)),
		H: int64(int(total) / (60 * 60) % 24),
	}
	countDown.F = getTimeString(countDown)
	return countDown
}

func StartCountDown(from, to time.Time) chan Countdown {
	total := to.Sub(from).Seconds()
	remaining := make(chan Countdown)
	go func() {
		for range time.Tick(1 * time.Second) {
			remainingTime := getTimeRemaining(to)
			remainingTime.P = 1 - (remainingTime.T / total)
			remaining <- remainingTime
		}
	}()
	return remaining
}

func FromIncomingContext(ctx context.Context, key string) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if authHeader, ok := md[key]; ok && len(authHeader) > 0 {
			return authHeader[0]
		}
	}
	return ""
}

func ConnectService(host string) (*grpc.ClientConn, error) {
	interceptor := grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req any, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if md, found := metadata.FromIncomingContext(ctx); found {
			ctx = metadata.NewOutgoingContext(ctx, md)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	})
	return grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()), interceptor)
}

func OutgoingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ignore := []string{}
	if !slices.Contains(ignore, info.FullMethod) {
		headers := []string{}
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			for key := range md {
				headers = append(headers, key, FromIncomingContext(ctx, key))
			}
		}
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(headers...))
	}
	return handler(ctx, req)
}

func UploadFile(dest, filename string, file []byte) (string, error) {
	resp, err := resty.New().R().SetFileReader("file", filename, bytes.NewReader(file)).Put(dest)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(resp.Body()), nil
}

func ReadMultipartFile(fileHeader *multipart.FileHeader) ([]byte, error) {
	if file, err := fileHeader.Open(); err != nil {
		return nil, errors.WithStack(err)
	} else {
		defer file.Close()
		return io.ReadAll(file)
	}
}

func RandInt(min, max int64) (int64, error) {
	if min > max {
		return 0, errors.New("min must be less than or equal to max")
	}
	if min == max {
		return min, nil
	}
	var buf [8]byte
	if _, err := rand.Read(buf[:]); err != nil {
		return 0, fmt.Errorf("failed to generate random number: %v", err)
	}
	randVal := binary.BigEndian.Uint64(buf[:])
	rangeSize := uint64(max - min + 1)
	return min + int64(randVal%rangeSize), nil
}

// FairRandomFloat generates a cryptographically secure random float64 in [min, max)
// Uses crypto/rand for true randomness suitable for gambling applications
func RandFloat(min, max float64) (float64, error) {
	if min >= max {
		return 0, fmt.Errorf("invalid range: min must be less than max")
	}

	// Generate a cryptographically secure random uint64
	var buf [8]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		return 0, fmt.Errorf("failed to generate random number: %v", err)
	}

	// Convert to uint64
	randUint := binary.BigEndian.Uint64(buf[:])

	// The maximum value for a uint64 (for normalization)
	const maxUint64 = float64(^uint64(0))

	// Normalize to [0,1) and scale to desired range
	normalized := float64(randUint) / (maxUint64 + 1.0) // +1 to ensure we never get 1.0
	scaled := min + normalized*(max-min)

	return scaled, nil
}

func CalculateLisenseExpiration(currentExp time.Time, subscription string, duration int64) int64 {
	switch subscription {
	case "monthly":
		return currentExp.Add(time.Duration(duration) * time.Hour * 24 * 30).Unix()
	case "yearly":
		return currentExp.Add(time.Duration(duration) * time.Hour * 24 * 365).Unix()
	}
	return currentExp.Add(time.Duration(duration) * time.Hour * 24 * 7).Unix()
}
