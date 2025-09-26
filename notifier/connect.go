package notifier

import (
	"github.com/thedivinez/go-libs/utils"
)

func Connect(addr string) (NotifierClient, error) {
	conn, err := utils.ConnectService(addr)
	if err != nil {
		return nil, err
	}
	return NewNotifierClient(conn), nil
}
