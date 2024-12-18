package services

import (
	"fmt"
	"log"
	"net"

	"github.com/thedivinez/go-libs/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Service struct {
	port     string
	Server   *grpc.Server
	listener net.Listener
}

func NewService(port string) (*Service, error) {
	if listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port)); err != nil {
		return nil, err
	} else {
		srv := grpc.NewServer(grpc.UnaryInterceptor(utils.OutgoingInterceptor))
		return &Service{port: port, Server: srv, listener: listener}, nil
	}
}

func (service *Service) Start() error {
	reflection.Register(service.Server)
	log.Printf("starting service on port:%s", service.port)
	return service.Server.Serve(service.listener)
}
