package main

import (
	"fmt"
)

type employee struct {
	id int
	name string	
	dept string
}

func main() {
	var coder employee
	coder.id = 1
	coder.name = "John Doe"
	coder.dept = "Engineering"

	banker := employee{2, "Jane Smith", "Finance"}	
	fmt.Println("Display coder", coder)
	fmt.Println("Display banker", banker)
	fmt.Printf("Coder ID: %d, Name: %s, Department: %s\n", coder.id, coder.name, coder.dept)

}