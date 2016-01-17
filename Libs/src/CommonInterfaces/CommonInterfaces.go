// commonInterfaces project commonInterfaces.go
// The package contains commenly used interfaces
package commonInterfaces

import (
	. "customErrors"
)

// Interface represents Initialize methood
type Initializer interface {
	// The method sets up instance and returns error if instance won't be initialized
	Initialize() *Error
}

// Interface represents MustInitialize methood
type MustInitializer interface {
	// The method sets up instance and panics if instance won't be initialized
	MustInitialize()
}

// Interface represents Dispose method
type Disposer interface {
	// The method "cleans" all internal referencies to let current instance to be garbage collected
	Dispose()
}
