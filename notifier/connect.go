package notifier

import (
	context "context"
	"encoding/json"

	"github.com/thedivinez/go-libs/storage"
	"github.com/thedivinez/go-libs/utils"
)

type Client struct {
	Client NotifierClient
	Redis  *storage.RedisCache
}

func NewClient(host string) *Client {
	return &Client{Redis: storage.NewRedisCache(host)}
}

func (client *Client) Connect(addr string) error {
	conn, err := utils.ConnectService(addr)
	if err != nil {
		return err
	}
	client.Client = NewNotifierClient(conn)
	return nil
}

func (client *Client) Listen(channels ...string) <-chan *EventMessage {
	message := make(chan *EventMessage)
	go func() {
		for msg := range client.Redis.Client.Subscribe(context.Background(), channels...).Channel() {
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
	Target  string `json:"target" bson:"target"`
	Message any    `json:"message" bson:"message"`
	Service string `json:"service" bson:"service"`
}

func (client *Client) Send(message *EventMessage) error {
	return client.Redis.Client.Publish(context.Background(), message.Target, message).Err()
}

func (ev *EventMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(ev)
}

func (ev *EventMessage) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, ev)
}
