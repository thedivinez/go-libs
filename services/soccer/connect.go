package soccer

import "github.com/thedivinez/go-libs/utils"

func Connect(addr string) (SoccerSimulatorClient, error) {
	conn, err := utils.ConnectService(addr)
	if err != nil {
		return nil, err
	}
	return NewSoccerSimulatorClient(conn), nil
}
