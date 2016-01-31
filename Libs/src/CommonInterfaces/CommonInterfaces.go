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

// The interface represents general Storage functionality
type IStorage interface {
	// Fills up model m with data from the first entry at Storage by:
	// if atleast one of PK field is not nil - then by PK fields
	// else by BK fields even if they are all nill
	// NB: the real type of m should be registered upon the operation. Otherwise returns Notsupported *Error
	Get(m interface{}) *Error

	// Stores data from model m into storage. Does not resolve any values by its own
	// NB: the real type of m should be registered upon the operation. Otherwise returns Notsupported *Error
	Put(m interface{}) *Error

	// Tries to store data from model m into storage
	// If there are fields which are resolved on storage level(sequence, default values or calculated)
	// fills up model with this values
	// NB: the real type of m should be registered upon the operation. Otherwise returns Notsupported *Error
	Resolve(m interface{}) *Error

	// Removes entry from storage by:
	// if at least one of PK field is not nil - then by PK fields
	// else by BK fields even if they are all nill
	// NB: the real type of m should be registered upon the operation. Otherwise returns Notsupported *Error
	Del(m interface{}) *Error
}
