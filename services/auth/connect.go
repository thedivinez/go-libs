package auth

import "github.com/thedivinez/go-libs/utils"

func Connect(addr string) (AuthenticationClient, error) {
	conn, err := utils.ConnectService(addr)
	if err != nil {
		return nil, err
	}
	return NewAuthenticationClient(conn), nil
}
