package broadcaster

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	redis *redis.Client
}

func NewClient(host string, db int) (*Client, error) {
	redisClient := redis.NewClient(&redis.Options{Addr: host, ReadTimeout: -1, DB: db})
	if err := redisClient.Ping(context.TODO()).Err(); err != nil {
		if err := redisClient.Ping(context.TODO()).Err(); err != nil {
			return nil, fmt.Errorf("failed to connect to messaging service")
		}
	}
	return &Client{redis: redisClient}, nil
}

func (client *Client) Listen(channels ...string) <-chan *EventMessage {
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

type EventMessage struct {
	Room    string `json:"room" bson:"room"`
	Event   string `json:"event" bson:"event"`
	OrgId   string `json:"orgId" bson:"orgId"`
	Message any    `json:"message" bson:"message"`
	Service string `json:"service" bson:"service"`
}

func (client *Client) Send(message *EventMessage) error {
	return client.redis.Publish(context.Background(), message.OrgId, message).Err()
}

func (ev *EventMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(ev)
}

func (ev *EventMessage) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, ev)
}
