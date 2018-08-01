package detailub

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

	ErrEmptyIdUser = errEmptyIdUser{}
	ErrEmptyIdBook = errEmptyIdBook{}
	ErrBookHanded  = errBookHanded{}
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

type errEmptyIdUser struct{}

func (errEmptyIdUser) Error() string {
	return "id User  is required"
}

func (errEmptyIdUser) StatusCode() int {
	return http.StatusBadRequest
}

type errEmptyIdBook struct{}

func (errEmptyIdBook) Error() string {
	return "Id Book  is required"
}

func (errEmptyIdBook) StatusCode() int {
	return http.StatusBadRequest
}

type errBookHanded struct{}

func (errBookHanded) Error() string {
	return "The book has been borrowed"
}

func (errBookHanded) StatusCode() int {
	return http.StatusBadRequest
}
