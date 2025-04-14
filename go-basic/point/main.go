package main

import (
	"fmt"
)

func main() {
	var num int = 10
	var numPtr *int = &num                       // Pointer to num
	var numPtr2 **int = &numPtr                  // Pointer to pointer to num
	fmt.Println("Value of num:", num)            // Output: 10
	fmt.Println("Value of numPtr:", *numPtr)     // Output: 10
	fmt.Println("Value of numPtr2:", **numPtr2)  // Output: 10
	fmt.Println("Address of num:", &num)         // Output: Address of num
	fmt.Println("Address of numPtr:", &numPtr)   // Output: Address of numPtr
	fmt.Println("Address of numPtr2:", &numPtr2) // Output: Address of numPtr2

	arrayPointers()
	structPointers()
	// 3. Pointer with a Slice
	slice := []int{4, 5, 6}
	fmt.Println("----- updateSlice with pointer ----")
	fmt.Println("Original Slice:", slice) // Output: [4, 5, 6]
	updateSlice(&slice)
	fmt.Println("Updated Slice:", slice) // Output: [40, 50, 60]
}

func arrayPointers() {
	fmt.Println("----- Array Pointers----")
	var arr = [3]int{1, 2, 3}
	fmt.Print(arr) // Output: [1 2 3]
	// Pointer to the array
	var p *[3]int = &arr
	(*p)[0] = 10     // Change the first element of arr through the pointer
	fmt.Println(arr) // Output: [10 2 3]
}

func structPointers() {
	fmt.Println("----- Struct Pointers----")
	type Person struct {
		Name string
		Age  int
	}
	var p = Person{"Alice", 25}
	fmt.Println("Before Person:", p) // Output: {Alice 25}
	var ptr *Person = &p
	ptr.Age = 30                            // Updates the Age field of the struct
	fmt.Println("After Person:", p)         // Output: {Alice 30}
	fmt.Println("Pointer to Person:", ptr)  // Output: &{Alice 30}
	fmt.Println("Pointer to Person:", *ptr) // Output: {Alice 30}
}

// Function to update a slice using a pointer
func updateSlice(s *[]int) {
	for i := range *s {
		(*s)[i] *= 10 // Multiply each element by 10
	}
}
