package quicksoccer

import (
	"github.com/thedivinez/go-libs/utils"
)

func Connect(addr string) (QuickSoccerClient, error) {
	conn, err := utils.ConnectService(addr)
	if err != nil {
		return nil, err
	}
	return NewQuickSoccerClient(conn), nil
}
