package WorkerPool

import (
	"errors"
	"time"
)

const (
	ERR_DEPARTMENTSHUTDOWN = "The department is shutting down and wont take new assignments"
)

var (
	sleepBetweenAttempts = time.Duration(1)
)

// Represents simple workers pool operating on Projects and WorkITems
type Department interface {

	// Synchronously requests worker slot and end exectes/does WorkItem in parallel routine as soon as slot is obtained
	Do(wi WorkItem) error

	// Returns true if there are taken slot
	IsBusy() bool

	// Returns true if no free slot available
	IsCompletelyBusy() bool

	// Closes the department
	Close()

	// Assigns a new Project to Depatment with deadline. prjDeadLine = 0 means no deadline for the project.
	// AssignNewProject(prj Project, prjDeadLine time.Duration, result <-chan (error)) error

	// Increeses number of workers for n
	// EmployWorkers(n int) error

	// Decreases number of workers for n
	// FireWorkers(n int) error
}

// Private custom  implementation of department
type depatment struct {
	numberOfWorkers    int
	isShuttingDown     bool
	workersExecutePool chan struct{}
	workersRequestPool chan struct{}
}

// A new Deapertment Factory
func NewDepartment(initWorkerNumber int) Department {
	// instantiate request pool
	workersRequestPool := make(chan struct{}, initWorkerNumber)

	// fill up request pool
	// for each initially empty slot we shoul put one value
	for i := 0; i < initWorkerNumber; i++ {
		workersRequestPool <- struct{}{}
	}

	return &depatment{
		numberOfWorkers:    initWorkerNumber,
		workersExecutePool: make(chan struct{}, initWorkerNumber),
		workersRequestPool: workersRequestPool,
	}

}

// Implements Department.Do(wi WorkItem) method
func (this *depatment) Do(wi WorkItem) error {
	reqC := this.requestSlot()
	_, more := <-reqC

	if this.isShuttingDown || !more {
		return errors.New(ERR_DEPARTMENTSHUTDOWN)
	}

	defer this.releaseSlot()

	this.workersExecutePool <- struct{}{}

	go func() {
		defer func() { <-this.workersExecutePool }()
		wi()
	}()

	return nil
}

func (this *depatment) IsBusy() bool {
	return len(this.workersExecutePool) > 0 || len(this.workersRequestPool) < this.numberOfWorkers
}

func (this *depatment) IsCompletelyBusy() bool {
	return len(this.workersRequestPool) == 0 || len(this.workersExecutePool) == this.numberOfWorkers
}

func (this *depatment) Close() {
	this.isShuttingDown = true

	// wait while all left assignments are done
	for len(this.workersExecutePool) > 0 || len(this.workersRequestPool) < this.numberOfWorkers {

	}

	close(this.workersRequestPool)

	if this.isShuttingDown {
		time.Sleep(2 * sleepBetweenAttempts)
	}
	close(this.workersExecutePool)

}

func (this *depatment) requestSlot() <-chan struct{} {
	return this.workersRequestPool
}

func (this *depatment) releaseSlot() {
	if !this.isShuttingDown {
		this.workersRequestPool <- struct{}{}
	}
}
