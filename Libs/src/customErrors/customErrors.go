// Errors project Errors.go
// The package contains general error exteded functionality
package customErrors

import (
	"fmt"
)

//go:generate stringer -type=ErrorType
type ErrorType int

const (
	BasicError ErrorType = iota
	InvalidOperation
	InvalidArgument
	AccessViolation
	Nonsupported
)

// Represents custom error

type Error struct {
	Type    ErrorType
	Message string
}

// Implementation of error interface
func (e Error) Error() string {
	return fmt.Sprintf("%T{Type:%s, Message:%s}", e, e.Type, e.Message)
}

// Error factory
func NewError(typ ErrorType, msg string) *Error {
	return &Error{
		Type:    typ,
		Message: msg,
	}
}
