package notifier

import (
	context "context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/thedivinez/go-libs/utils"
)

type Client struct {
	redis *redis.Client
}

func NewClient(host string) *Client {
	opts, err := redis.ParseURL(host)
	if err != nil {
		log.Fatalf("Error parsing Redis URL: %v", err)
	}
	opts.ReadTimeout = -1
	return &Client{redis: redis.NewClient(opts)}
}

func (client *Client) Connect(addr string) (NotifierClient, error) {
	conn, err := utils.ConnectService(addr)
	if err != nil {
		return nil, err
	}
	return NewNotifierClient(conn), nil
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
