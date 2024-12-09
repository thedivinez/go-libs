package messaging

import "encoding/json"

type EventMessage struct {
	Message any
	Room    string
	Event   string
	Service string
	OrgId   string
}

func (ev EventMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(ev)
}

func (ev *EventMessage) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, ev)
}
