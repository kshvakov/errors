package errors

import (
	"testing"
)

func Err(message string) error {

	return Newf(message, "err", 42)
}

func TestExample(t *testing.T) {

	err := Err("%s: message (%d)")

	if cErr, ok := err.(Error); ok {

		t.Log(cErr.Error(), cErr.Stack())

	} else {

		t.Log(err.Error())
	}
}
