package rps

import "encoding/json"

func (s *RPSBet) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *RPSBet) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
