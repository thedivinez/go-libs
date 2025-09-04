package skyrush

import "github.com/thedivinez/go-libs/utils"

func Connect(addr string) (SkyRushClient, error) {
	conn, err := utils.ConnectService(addr)
	if err != nil {
		return nil, err
	}
	return NewSkyRushClient(conn), nil
}
