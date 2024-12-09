package socketio

import (
	"github.com/zishang520/socket.io/v2/socket"
)

type Client struct {
	*socket.Socket
	UserId   string
	DeviceId string
}

type Server struct {
	*socket.Server
}

type AckFunc func(response any)

type Event struct {
	Name string
	Data any
	Ack  AckFunc
}

func Room(room string) socket.Room {
	return socket.Room(room)
}

func NewServer(srv any, opts socket.ServerOptionsInterface) *Server {
	return &Server{socket.NewServer(srv, opts)}
}

func OnConnect(connect func(client *Client)) func(clients ...any) {
	return func(clients ...any) {
		client := clients[0].(*socket.Socket)
		userId, _ := client.Request().Headers().Get("X-User-Id")
		deviceId, _ := client.Request().Headers().Get("X-Device-Id")
		connect(&Client{client, userId, deviceId})
	}
}

func handleAck(handler func([]any, error)) AckFunc {
	return func(data any) {
		if handler != nil {
			handler([]any{data}, nil)
		}
	}
}

func (client *Client) On(event string, handler func(client *Client, msg *Event)) {
	client.Socket.On(event, func(data ...any) {
		message := any(nil)
		ack := func([]any, error) {}
		if len(data) > 1 {
			if _, ok := data[1].(func([]any, error)); ok {
				ack = data[1].(func([]any, error))
			}
		}
		if len(data) > 0 {
			message = data[0]
		}
		handler(client, &Event{event, message, handleAck(ack)})
	})
}

func (client *Client) To(room ...socket.Room) *socket.BroadcastOperator {
	return client.Socket.To(room...)
}
