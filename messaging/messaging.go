package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Messenger struct {
	redis *redis.Client
}

func NewClient(host string, db int) (*Messenger, error) {
	redisClient := redis.NewClient(&redis.Options{Addr: host, ReadTimeout: -1, DB: db})
	if err := redisClient.Ping(context.TODO()).Err(); err != nil {
		time.Sleep(3 * time.Second)
		if err := redisClient.Ping(context.TODO()).Err(); err != nil {
			return nil, fmt.Errorf("failed to connect to messaging service")
		}
	}
	return &Messenger{redis: redisClient}, nil
}

func (client *Messenger) Listen(channels ...string) <-chan *EventMessage {
	message := make(chan *EventMessage)
	go func() {
		for msg := range client.redis.Subscribe(context.Background(), channels...).Channel() {
			payload := &EventMessage{}
			if err := payload.UnmarshalBinary([]byte(msg.Payload)); err == nil {
				message <- payload
			}
		}
		close(message)
	}()
	return message
}

func (client *Messenger) Send(message EventMessage) error {
	return client.redis.Publish(context.Background(), message.OrgId, message).Err()
}
