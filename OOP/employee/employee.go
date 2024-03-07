package employee

import (
	"fmt"
)

type employee struct {  
    firstName   string
    lastName    string
    totalLeaves int
    leavesTaken int
}

func New(firstName string, lastName string, totalLeaves int, leavesTaken int) employee {
	e := employee {
		firstName,
		lastName,
		totalLeaves,
		leavesTaken,
	}
	return e
}

func (e employee) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining\n", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}

// since employee is unexported, it’s not possible to create values of type Employee from other packages. Hence we are providing an exported New function which takes the required parameters as input and returns a newly created employee.

// we have created an employee struct and attached a method named leavesremaining()
// the struct and method together is akined to a class