package rps

import "encoding/json"

func (s *Bet) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Bet) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
