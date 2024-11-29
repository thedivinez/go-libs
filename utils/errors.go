package utils

import (
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
)

type ServiceError struct {
	Code    codes.Code `json:"code"`
	Message string     `json:"message"`
}

func NewServiceError(code codes.Code, message string) *ServiceError {
	return &ServiceError{Code: code, Message: message}
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf(`{"code":%d,"message":"%s"}`, e.Code, e.Message)
}

func (e *ServiceError) WithInternal(err error) *ServiceError {
	log.Println(err)
	return e
}
