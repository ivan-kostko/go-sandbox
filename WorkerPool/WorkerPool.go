package WorkerPool

// Represents simple assignment
type WorkItem func()

// Represents collection of WorkItems with their dependencies, deadline, and failure delegate
type Project struct {
	WorkItems []WorkItem
}

type Enterprise interface {

	// Creates a new department with initWorkerNumber workers available
	AddNewDepartment(name string, initWorkerNumber int) error

	// Moves n workers from original to dest department
	MoveWorkersBetweenDepartments(n int, original, dest string) error

	// Represents number of all workes over all departments
}
