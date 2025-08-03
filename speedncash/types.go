package speedncash

import "encoding/json"

func (s *RaceBet) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *RaceBet) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
