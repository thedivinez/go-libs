package rps

import "github.com/thedivinez/go-libs/utils"

func Connect(addr string) (RockPaperScissorsClient, error) {
	conn, err := utils.ConnectService(addr)
	if err != nil {
		return nil, err
	}
	return NewRockPaperScissorsClient(conn), nil
}
