// Errors project Errors.go
// The package contains general error exteded functionality
package customErrors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestErrorImplements_errorInterface(t *testing.T) {
	expectedInterface := reflect.TypeOf((*error)(nil)).Elem()
	err := Error{}
	if !reflect.TypeOf(err).Implements(expectedInterface) {
		t.Errorf("%T does not implement %v", err, expectedInterface)
	}
}

func TestErrorErrorMethod(t *testing.T) {
	expected := `customErrors.Error{Type:AccessViolation, Message:Testing}`
	var err error = Error{Type: AccessViolation, Message: "Testing"}
	if err.Error() != expected {
		t.Errorf("err.Error() returned %v while expected %v", err.Error(), expected)
	}
}

func TestErrorTypeStringer(t *testing.T) {
	expected := "AccessViolation"
	if fmt.Sprintf("%s", AccessViolation) != expected {
		t.Errorf("Returned %s while expected %v", AccessViolation, expected)
	}
}
