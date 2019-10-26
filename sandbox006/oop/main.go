package main

import (
	"github.com/codegold79/cai/sandbox006/oop/employee"
)

func main() {
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
