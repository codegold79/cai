// Structs Instead of Classes tutorial from
// https://golangbot.com/structs-instead-of-classes/

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
	e := employee{firstName, lastName, totalLeaves, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("\n%s %s has %d leaves remaining.\n", e.firstName, e.lastName, e.totalLeaves-e.leavesTaken)
}
