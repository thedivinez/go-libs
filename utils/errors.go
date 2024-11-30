package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

type ServiceError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewServiceError(code int, message string) *ServiceError {
	return &ServiceError{Code: code, Message: message}
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf(`{"code":%d,"message":"%s"}`, e.Code, e.Message)
}

func (e *ServiceError) WithInternal(err error) *ServiceError {
	log.Println(err)
	return e
}

func FromServiceError(err error) (*ServiceError, error) {
	errorMessage := &ServiceError{Code: 500, Message: "internal error"}
	if rpcErr, ok := status.FromError(err); !ok {
		return errorMessage, errors.New(err.Error())
	} else {
		if err := json.Unmarshal([]byte(rpcErr.Message()), &errorMessage); err != nil {
			return errorMessage, err
		}
	}
	return errorMessage, nil
}
