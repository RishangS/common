package customerror

import "fmt"

const (
	DataRetrieveException       = "failed to retrieve data"
	QueueCreationErrorException = "failed to create queue"
	MarshalException            = "failed to marshal"
	ConsumerRegisterException   = "failed to register consumer"
)

type CustomError struct {
	errType string
	msg     string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.errType, e.msg)
}

// NewCustomError creates a new error with the provided type and optional additional message
func NewCustomError(errType string, errString ...string) error {
	if len(errString) > 0 {
		return &CustomError{errType, fmt.Sprintf("%s : %v", errType, errString)}
	}
	return &CustomError{errType, errType}
}

// Specific error creators
func NewDataRetrieveExceptionError(errString ...string) error {
	return NewCustomError(DataRetrieveException, errString...)
}

func NewQueueCreationErrorExceptionError(errString ...string) error {
	return NewCustomError(QueueCreationErrorException, errString...)
}

func NewMarshalExceptionError(errString ...string) error {
	return NewCustomError(MarshalException, errString...)
}

func NewConsumerRegisterExceptionError(errString ...string) error {
	return NewCustomError(ConsumerRegisterException, errString...)
}
