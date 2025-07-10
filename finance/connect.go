package finance

import "github.com/thedivinez/go-libs/utils"

func Connect(addr string) (PaymentClient, error) {
	conn, err := utils.ConnectService(addr)
	if err != nil {
		return nil, err
	}
	return NewPaymentClient(conn), nil
}
