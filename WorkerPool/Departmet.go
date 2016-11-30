package WorkerPool

import (
	"errors"
	"time"
)

const (
	ERR_DEPARTMENTSHUTDOWN = "The department is shutting down and wont take new assignments"
	ERR_TIMEDOUTREQUSTSLOT = "The request for a free execution slot has been timed out"
)

// Represents simple workers pool operating on Projects and WorkITems
type Department interface {

	// Synchronously requests worker slot and end exectes/does WorkItem in parallel routine as soon as slot is obtained
	// If no slot aquired upon timeOut exceeds - returns ERR_TIMEDOUTREQUSTSLOT
	Do(wi WorkItem, timeOut time.Duration) error

	// Closes the department
	Close()
}

// Private custom  implementation of department
type depatment struct {
	isShuttingDown bool
	workersPool    chan struct{}
}

// A new Deapertment Factory
func NewDepartment(initWorkerNumber int) Department {
	// instantiate  pool
	workersPool := make(chan struct{}, initWorkerNumber)

	// fill up pool
	// for each initially empty slot we shoul put one value
	for i := 0; i < initWorkerNumber; i++ {
		workersPool <- struct{}{}
	}

	return &depatment{
		isShuttingDown: false,
		workersPool:    workersPool,
	}

}

// Implements Department.Do(wi WorkItem) method
func (this *depatment) Do(wi WorkItem, timeOut time.Duration) error {

	if this.isShuttingDown {
		return errors.New(ERR_DEPARTMENTSHUTDOWN)
	}

	t := time.NewTimer(timeOut)

	select {
	case _ = <-this.workersPool:
		if !t.Stop() {
			<-t.C
		}
	case _ = <-t.C:
		return errors.New(ERR_TIMEDOUTREQUSTSLOT)
	}

	if this.isShuttingDown {
		return errors.New(ERR_DEPARTMENTSHUTDOWN)
	}

	go func() {
		defer this.releaseSlot()
		wi()
	}()

	return nil
}

func (this *depatment) Close() {
	this.isShuttingDown = true

	// wait while all left assignments are done
	for i := 0; i < cap(this.workersPool); i++ {
		<-this.workersPool
	}

	close(this.workersPool)

}

func (this *depatment) releaseSlot() {
	this.workersPool <- struct{}{}
}
