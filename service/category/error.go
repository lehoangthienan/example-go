package category

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound       = errNotFound{}
	ErrUnknown        = errUnknown{}
	ErrNameIsRequired = errNameIsRequired{}
	ErrRecordNotFound = errRecordNotFound{}

	ErrLenghtName  = errLenghtName{}
	ErrNameExisted = errNameExisted{}
)

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}
func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errUnknown struct{}

func (errUnknown) Error() string {
	return "unknown error"
}
func (errUnknown) StatusCode() int {
	return http.StatusBadRequest
}

type errRecordNotFound struct{}

func (errRecordNotFound) Error() string {
	return "client record not found"
}
func (errRecordNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errNameIsRequired struct{}

func (errNameIsRequired) Error() string {
	return "user name is required"
}

func (errNameIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errLenghtName struct{}

func (errLenghtName) Error() string {
	return "Name is not shorter five character"
}

func (errLenghtName) StatusCode() int {
	return http.StatusBadRequest
}

type errNameExisted struct{}

func (errNameExisted) Error() string {
	return "Name is not existed"
}

func (errNameExisted) StatusCode() int {
	return http.StatusBadRequest
}
