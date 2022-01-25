package database

import (
	"fmt"
)

var (
	//warn
	WarnNoRowsAffected = NewError("no rows were affected")

	//error
	ErrQueryRow           = NewError("error query row")
	ErrBeginTransaction   = NewError("error when starting transaction")
	ErrRowsAffected       = NewError("error getting rows affected")
	ErrLastInsertID       = NewError("error lastInsertID may not be supported")
	ErrQuery              = NewError("error when executing query")
)

type Error struct {
	detail  string
	cause   error
	subject string
}

func (e Error) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("error: %s, caused by: %s", e.detail, e.cause.Error())
	}

	return e.detail
}

func NewError(d string) Error {
	return Error{detail: d}
}

func (e Error) WithCause(err error) Error {
	e.cause = err
	return e
}

func (e Error) Is(err error) bool {
	return e.detail == err.Error()
}

func (e Error) WithSubject(subject string) Error {
	e.subject = subject
	return e
}

func (e Error) Subject() string {
	return e.subject
}