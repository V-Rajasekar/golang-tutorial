// Package main demonstrates various Go programming concepts including:
// - Function definitions and usage
// - Passing arguments by reference
// - Using switch statements for conditional logic
// - Passing functions as arguments
// - Defining and using custom function types
package main

import "fmt"


func main() {
	var nights int = 3
	var rate float32 = 145.90
	fmt.Println(computePrice(rate, nights))
	passReference()
	passFunctions()

}

// compute the price with a 200% margin
func computePrice(rate float32, nights int) (price float32) {
	p := rate * float32(nights)
	p = p * 2
	return
}

func occupancyLevel(occupancyRate float32) string {
	switch {
	case occupancyRate > 70:
		return "High"
	case occupancyRate > 20:
		return "Medium"
	default:
		return "Low"
	}
}

// passReference demonstrates passing arguments by reference in Go.
// It swaps the values of two integers using pointers.
func passReference() {
	// Pass by reference
	var a, b int = 1, 2
	fmt.Println("Before swap:", a, b)
	swap(&a, &b)
	fmt.Println("After swap:", a, b)
}

// swap exchanges the values of two integers using their memory addresses.
// Parameters:
// - x: A pointer to the first integer.
// - y: A pointer to the second integer.
func swap(x, y *int) {
	// Swap the values of x and y
	temp := *x
	*x = *y
	*y = temp
}

// Pass functions as arguments

type adder func(a, b int) int

var add adder = func(a, b int) int {
	return a + b
}

// dub is a higher-order function that takes another function as an argument,
func dub(twice func(int, int) int) {
	fmt.Println("Twice:", twice(2, 3)*2) // Output: Twice: 10
}

// passFunctions demonstrates passing functions as arguments in Go.
// It uses a custom function type and passes a function to another function.
func passFunctions() {
	dub(add) // Pass the add function as an argument
}
