package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/thedivinez/go-libs/utils"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(address string, db int) *RedisCache {
	redisClient := redis.NewClient(&redis.Options{Addr: address, ReadTimeout: -1, DB: db})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		time.Sleep(3 * time.Second)
		if err := redisClient.Ping(context.Background()).Err(); err != nil {
			log.Fatal("failed to connect to redis :", err)
		}
	}
	return &RedisCache{Client: redisClient}
}

func (client *RedisCache) transcode(in, out any) error {
	resultBytes, err := json.Marshal(in)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(json.Unmarshal(resultBytes, &out))
}

func (client *RedisCache) Write(key, path string, value interface{}) error {
	if reflect.TypeOf(value).String() == "string" {
		value = fmt.Sprintf(`"%s"`, value)
	}
	return errors.WithStack(client.Client.JSONSet(context.Background(), key, path, value).Err())
}

func (client *RedisCache) Read(key, path string, dest interface{}) error {
	if res, err := client.Client.JSONGet(context.Background(), key, path).Expanded(); err == nil {
		result, ok := res.([]any)
		if !ok {
			return errors.New("document not found in cache")
		}
		if len(result) > 0 {
			if strings.Contains(reflect.TypeOf(dest).String(), "[]") {
				if strings.Contains(reflect.TypeOf(result[0]).String(), "[]") {
					return client.transcode(result[0], dest)
				}
				return client.transcode(result, dest)
			}
			return client.transcode(result[0], dest)
		}
		if strings.Contains(reflect.TypeOf(dest).String(), "[]") {
			return nil
		}
		return errors.New("document not found in cache")
	} else {
		return errors.WithStack(err)
	}
}

func (client *RedisCache) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanIterator {
	return client.Client.Scan(ctx, cursor, match, count).Iterator()
}

func (client *RedisCache) Remove(key string) error {
	return client.Client.Del(context.Background(), key).Err()
}

func (client *RedisCache) Count(key string) int64 {
	if result, err := client.Client.JSONArrLen(context.Background(), key, "$").Result(); err != nil {
		log.Println(err)
	} else {
		if len(result) > 0 {
			return result[0]
		}
	}
	return 0
}

func (client *RedisCache) Publish(userType string, message utils.EventMessage) error {
	message.Room = fmt.Sprintf("%s:%s", userType, message.Room)
	return client.Client.Publish(context.Background(), message.Service, message).Err()
}
