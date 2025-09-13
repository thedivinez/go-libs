package storage

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	DbAddress string `json:"DB_ADDRESS"`
	DbName    string `json:"DB_NAME"`
}

type MongoStorage struct{ db *mongo.Database }

func NewMongoStorage(config MongoDBConfig) *MongoStorage {
	clientOption := options.Client().ApplyURI(config.DbAddress)
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	return &MongoStorage{db: client.Database(config.DbName)}
}

func (s *MongoStorage) Transcode(in, out any) error {
	resultBytes, err := json.Marshal(in)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(json.Unmarshal(resultBytes, &out))
}

/*******************************************
*	FIND ALL DOCUMENTS IN COLLECTION METHOD  *
*******************************************/
func (s *MongoStorage) Find(collection string, filter any, results interface{}, opts ...*options.FindOptions) error {
	res, dbError := s.db.Collection(collection).Find(context.TODO(), filter, opts...)
	if dbError != nil {
		return errors.WithStack(dbError)
	}
	defer res.Close(context.TODO())
	return errors.WithStack(res.All(context.TODO(), results))
}

/*****************************************
*	FIND ONE DOCUMENT IN COLLECTION	METHOD *
*****************************************/
func (s *MongoStorage) FindOne(collection string, filter any, results interface{}, opts ...*options.FindOneOptions) error {
	return errors.WithStack(s.db.Collection(collection).FindOne(context.TODO(), filter, opts...).Decode(results))
}

/*******************************************
*	DELETE ONE DOCUMENT IN COLLECTION	METHOD *
*******************************************/
func (s *MongoStorage) DeleteOne(collection string, filter any) error {
	_, err := s.db.Collection(collection).DeleteOne(context.TODO(), filter)
	return errors.WithStack(err)
}

/*******************************************
*	DELETE ONE DOCUMENT IN COLLECTION	METHOD *
*******************************************/
func (s *MongoStorage) DeleteMany(collection string, filter any) error {
	_, err := s.db.Collection(collection).DeleteMany(context.TODO(), filter)
	return errors.WithStack(err)
}

/********************************************
*	INSERT ONE DOCUMENT IN COLLECTION METHOD  *
********************************************/
func (s *MongoStorage) InsertOne(collection string, data any) (insertId string, err error) {
	result, err := s.db.Collection(collection).InsertOne(context.TODO(), data)
	if reflect.TypeOf(result.InsertedID).String() == "primitive.ObjectID" {
		insertId = result.InsertedID.(primitive.ObjectID).Hex()
	} else {
		insertId = result.InsertedID.(string)
	}
	return insertId, errors.WithStack(err)
}

func (s *MongoStorage) InsertMany(collection string, data ...any) (err error) {
	betsList := bson.A{}
	if err := s.Transcode(data[0], &betsList); err != nil {
		return errors.WithStack(err)
	}
	if len(betsList) > 0 {
		if _, err := s.db.Collection(collection).InsertMany(context.TODO(), betsList); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

/*******************************************
*	UPDATE ONE DOCUMENT IN COLLECTION	METHOD *
*******************************************/
func (s *MongoStorage) UpdateOne(collection string, filter any, target any, opts ...*options.UpdateOptions) error {
	if result, err := s.db.Collection(collection).UpdateOne(context.TODO(), filter, target, opts...); err != nil {
		return errors.WithStack(err)
	} else {
		if result.MatchedCount <= 0 {
			return errors.WithStack(errors.New("document not found"))
		}
		return nil
	}
}

/*******************************************
*	UPDATE Many DOCUMENT IN COLLECTION	METHOD *
*******************************************/
func (s *MongoStorage) UpdateMany(collection string, filter any, target any, opts ...*options.UpdateOptions) error {
	if _, err := s.db.Collection(collection).UpdateMany(context.TODO(), filter, target, opts...); err != nil {
		return errors.WithStack(err)
	} else {
		return errors.WithStack(err)
	}
}

/*******************************************
*	REPLACE ONE DOCUMENT IN COLLECTION	METHOD *
*******************************************/
func (s *MongoStorage) ReplaceOne(collection string, filter any, target any, opts ...*options.ReplaceOptions) error {
	if _, err := s.db.Collection(collection).ReplaceOne(context.TODO(), filter, target, opts...); err != nil {
		return errors.WithStack(err)
	} else {
		return errors.WithStack(err)
	}
}

/*******************************************
*	UPDATE ONE DOCUMENT IN COLLECTION	METHOD *
*******************************************/
func (s *MongoStorage) Count(collection string, filter any) (int64, error) {
	return s.db.Collection(collection).CountDocuments(context.TODO(), filter)
}

func (s *MongoStorage) GetPage(collection string, filter any, page string, limit, sort int64, results any) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	skip := int64(0)
	l := limit
	if page != "" {
		idx, err := strconv.Atoi(page)
		if err != nil {
			return 0, err
		}
		skip = int64(idx)*limit - limit
	}
	fOpt := options.FindOptions{Limit: &l, Skip: &skip, Sort: bson.M{"_id": sort}}
	res, err := s.db.Collection(collection).Find(ctx, filter, &fOpt)
	if err != nil {
		return 0, err
	}
	count, err := s.db.Collection(collection).CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	ceilDiv := func(a, b int64) int64 {
		if a%b == 0 {
			return a / b
		}
		return a/b + 1
	}
	return ceilDiv(count, limit), errors.WithStack(res.All(context.TODO(), results))
}

func (s *MongoStorage) Aggregate(collection string, filter any, results any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := s.db.Collection(collection).Aggregate(ctx, filter)
	if err != nil {
		return err
	}
	return errors.WithStack(res.All(context.TODO(), results))
}

func generateRandomNumber(numberOfDigits int) (int, error) {
	maxLimit := int64(int(math.Pow10(numberOfDigits)) - 1)
	lowLimit := int(math.Pow10(numberOfDigits - 1))
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxLimit))
	if err != nil {
		return 0, err
	}
	randomNumberInt := int(randomNumber.Int64())
	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}
	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}
	return randomNumberInt, nil
}

type ID struct {
	Value string
}

func (s *MongoStorage) GenerateID(collection string, letters, digits int) *ID {
	generate := func() (*ID, error) {
		letters := make([]byte, letters)
		letterBytes := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for i := range letters {
			randomNumber, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
			letters[i] = letterBytes[randomNumber.Int64()]
		}
		_digits, err := generateRandomNumber(digits)
		if err != nil {
			return nil, err
		}
		return &ID{Value: fmt.Sprintf("%s%06d", string(letters), _digits)}, nil
	}
	for range 10 {
		if id, err := generate(); err == nil {
			if err := s.FindOne(collection, bson.M{"id": id}, bson.M{}); err != nil {
				return id
			}
		}
	}
	return nil
}
