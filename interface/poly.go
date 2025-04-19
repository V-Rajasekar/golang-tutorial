package main

import "fmt"

type bird interface {
	speak() string
	move() string
}

type parrot struct {
}

type duck struct{}

func (p *parrot) speak() string {
	return "Parrot speaks"
}
func (p *parrot) move() string {
	return "Parrot flies"
}

func (d duck) speak() string {
	return "Duck quacks"
}
func (d duck) move() string {
	return "Duck swims"
}

func nudge(b bird) {
	fmt.Println(b.speak())
	fmt.Println(b.move())
}

func main() {
	b := &parrot{}
	nudge(b) // This will work as parrot implements the bird interface
	var duck duck
	nudge(duck)// This will work as duck implements the bird interface
}
