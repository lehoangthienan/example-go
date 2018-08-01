package book

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound        = errNotFound{}
	ErrUnknown         = errUnknown{}
	ErrNameIsRequired  = errNameIsRequired{}
	ErrEmailIsRequired = errEmailIsRequired{}
	ErrEmailIsInvalid  = errEmailIsInvalid{}
	ErrRecordNotFound  = errRecordNotFound{}

	ErrNameSake          = errNameSake{}
	ErrLenghtName        = errLenghtName{}
	ErrLenghtDescription = errLenghtDescription{}
	ErrIsDescription     = errIsDescription{}
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

type errEmailIsRequired struct{}

func (errEmailIsRequired) Error() string {
	return "email is required"
}
func (errEmailIsRequired) StatusCode() int {
	return http.StatusBadRequest
}

type errEmailIsInvalid struct{}

func (errEmailIsInvalid) Error() string {
	return "email address is invalid"
}
func (errEmailIsInvalid) StatusCode() int {
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

type errNameSake struct{}

func (errNameSake) Error() string {
	return "Name is Saked"
}

func (errNameSake) StatusCode() int {
	return http.StatusBadRequest
}

type errLenghtName struct{}

func (errLenghtName) Error() string {
	return "Name must >=5 character"
}

func (errLenghtName) StatusCode() int {
	return http.StatusBadRequest
}

type errLenghtDescription struct{}

func (errLenghtDescription) Error() string {
	return "Description must >= 5 character"
}

func (errLenghtDescription) StatusCode() int {
	return http.StatusBadRequest
}

type errIsDescription struct{}

func (errIsDescription) Error() string {
	return "Description is not null"
}

func (errIsDescription) StatusCode() int {
	return http.StatusBadRequest
}
