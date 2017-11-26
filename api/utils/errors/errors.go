package errors

import (
	"fmt"

	e "github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() e.StackTrace
}

func New(msg string) error {
	return e.New(msg)
}

func Wrap(err error, msg string) error {
	return e.Wrap(err, msg)
}

func Cause(err error) error {
	return e.Cause(err)
}

func StackTrace(err error) string {
	t, ok := Cause(err).(stackTracer)
	if !ok {
		return err.Error()
	}

	st := t.StackTrace()
	return fmt.Sprintf("%+v", st[1:5])
}
