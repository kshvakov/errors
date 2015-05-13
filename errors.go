package errors

import (
	"fmt"
)

func New(message string) error {

	return Raw(message, 3)
}

func Newf(format string, a ...interface{}) error {

	return Raw(fmt.Sprintf(format, a...), 3)
}

func Wrap(err error) error {

	return Raw(err.Error(), 3)
}

func Raw(message string, skip int) error {

	return &err{
		message: message,
		stack:   stack(skip),
	}
}

type err struct {
	message string
	stack   string
}

func (e *err) Error() string {

	return e.message
}

func (e *err) Stack() string {

	return e.stack
}
