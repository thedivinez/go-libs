package aviator

import "github.com/thedivinez/go-libs/utils"

func Connect(addr string) (AviatorClient, error) {
	conn, err := utils.ConnectService(addr)
	if err != nil {
		return nil, err
	}
	return NewAviatorClient(conn), nil
}
