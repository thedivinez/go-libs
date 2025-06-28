package aviator

import "encoding/json"

func (s *RoundBet) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *RoundBet) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
