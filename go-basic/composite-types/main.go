package main

import (
	"fmt"
	"slices"
)

func main() {
	array()
	slices_demo()
	slicingSlices()
	sliceCopy()
	convArraysToSlice()
	convSliceToArrays()
}

func array() {

	var x [3]int

	fmt.Println("x:", x)

	var y = [3]int{10, 20, 30}

	fmt.Printf("len(y):%v of elements:%v ", len(y), y)

}

func slices_demo() {
	var x = []int{10, 20, 30}
	var y = []int{10, 20, 30}
	fmt.Println("len(x): ", len(x))
	x[0] = 5

	var z []int // as no values assigned
	fmt.Println(z == nil)
	fmt.Printf("x: %v y:%v \n", x, y)
	fmt.Println("slice Equals:", slices.Equal(x, y))
	y = append(y, 40, 50, 60) // you can append more than one value

	//capacity

	fmt.Printf("y len: %v and capacity: %v \n", len(y), cap(y))
	y = append(y, 70)
	fmt.Printf("y len: %v and capacity: %v \n", len(y), cap(y))
	y = append(y, 80)
	fmt.Printf("y len: %v and capacity: %v \n", len(y), cap(y))
	//make

	a := make([]int, 5) // Empty slice using make of size 5
	fmt.Println("a:", a)
	b := make([]int, 5, 10) // Emplty slice with size and capacity
	fmt.Println("b:", b)

	fmt.Println("Emptying a Slice")
	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s))
	clear(s)
	fmt.Println(s, len(s))

	var nilSlice []int
	var emptySlice = []int{}
	fmt.Println(nilSlice == nil)
	fmt.Println(emptySlice == nil)
}

func slicingSlices() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}

func sliceCopy() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)
}

func convArraysToSlice() {
	xArray := [4]int{5, 6, 7, 8}
	xSlice := xArray[:]
	fmt.Printf("xSlice: %T \n", xSlice)
}

func convSliceToArrays() {
	xSlice := []int{1, 2, 3, 4} // Slice
	xArray := [2]int(xSlice)
	fmt.Printf("xArray: %T \n", xArray)
}
