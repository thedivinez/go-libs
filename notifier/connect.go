package notifier

import (
	"encoding/json"

	"github.com/thedivinez/go-libs/utils"
)

func Connect(addr string) (NotifierClient, error) {
	conn, err := utils.ConnectService(addr)
	if err != nil {
		return nil, err
	}
	return NewNotifierClient(conn), nil
}

func (ev *EventMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(ev)
}

func (ev *EventMessage) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, ev)
}
