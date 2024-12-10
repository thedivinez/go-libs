package aviator

import "encoding/json"

func (s *PlaneBet) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *PlaneBet) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
