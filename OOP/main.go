package main

import (
	"golangbot-course/OOP/employee"
)

func main() {
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}