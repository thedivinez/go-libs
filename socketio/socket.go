package socketio

import (
	"context"

	"github.com/zishang520/socket.io/v2/socket"
)

type SocketClient struct {
	context.Context
	*socket.Socket
	UserID    string
	DeviceID  string
	Container string
}

func NewSocketClient(connect func(client *SocketClient)) func(clients ...any) {
	return func(clients ...any) {
		client := clients[0].(*socket.Socket)
		ctx := context.WithoutCancel(client.Request().Context())
		userId, _ := client.Request().Headers().Get("X-User-Id")
		deviceID, _ := client.Request().Headers().Get("X-Device-Id")
		container, _ := client.Request().Headers().Get("X-Container")
		connect(&SocketClient{ctx, client, userId, deviceID, container})
	}
}

func (client *SocketClient) On(event string, handler func(client *SocketClient, msg []any)) {
	client.Socket.On(event, func(data ...any) {
		handler(client, data)
	})
}
