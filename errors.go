package errors

import (
	"fmt"
	"reflect"
)

type Error struct {
	Pos
	Err error
}

func (e *Error) Error() string {
	msg := fmt.Sprintf("(%s) %s", typeName(e.Err), e.Err.Error())
	return e.Pos.Decorate(msg, "", "")
}

func New(text string) error {
	return &Error{GetPos(1), errorString{text}}
}

func Format(format string, v ...interface{}) error {
	return &Error{GetPos(1), fmt.Errorf(format, v...)}
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	if _, isError := err.(*Error); isError {
		return err
	}
	return &Error{GetPos(1), err}
}

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func typeName(v interface{}) string {
	typ := reflect.TypeOf(v)
	pkg, name, str := typ.PkgPath(), typ.Name(), typ.String()
	if name == "" {
		return str
	}
	return pkg + "." + name
}
