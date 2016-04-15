// Errors project Errors.go
// The package contains general error exteded functionality
package CustomErrors

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

// Represents custom error as tuple Type + Message.
type Error struct {
	Type    ErrorType
	Message string
}

// Implementation of standart error interface
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
