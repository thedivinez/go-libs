package speedncash

import "github.com/thedivinez/go-libs/utils"

func Connect(addr string) (SpeedNCashClient, error) {
	conn, err := utils.ConnectService(addr)
	if err != nil {
		return nil, err
	}
	//
	return NewSpeedNCashClient(conn), nil
}
