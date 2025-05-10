// Package main demonstrates the use of embedded interfaces in Go.
// 
// Embedded interfaces allow you to compose multiple interfaces into a single interface.
// This is useful when you want to define a new interface that combines the behaviors
// of multiple existing interfaces.
// Allows you to create a further abstraction layer on top of existing interfaces.
// This can be useful for creating more complex types that need to implement multiple behaviors.
package main

import (
	"fmt"
)

type bird interface {
	speak() string	
	move() string
}

type human interface {
	speak() string		
	move() string
}

//embed interface
type creature interface {
	bird
	human
}

type parrot struct {
}

func (parrot) speak() string {
	return "Parrot speaks"
}
func (parrot) move() string {
	return "Parrot flies"
}	

type person struct {
}
func (person) speak() string {
	return "Person speaks"
}
func (person) move() string {
	return "Person walks"
}

func nudge(c creature) {
	fmt.Printf("\n%v\n", c.speak())
	fmt.Printf("%v\n", c.move())
}

func main() {
	var bird1 parrot
	var human1 person
	nudge(bird1) // This will work as parrot implements the bird interface
	nudge(human1) // This will work as person implements the human interface
}