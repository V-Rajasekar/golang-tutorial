package main

import (
	"fmt"
)

func main() {
	var cars [3]string
	fmt.Printf("Display arr1 len=%d cap=%d value=%v \n", len(cars), cap(cars), cars)

	cars[0] = "Toyota"
	cars[1] = "Honda"
	cars[2] = "Ford"
	fmt.Printf("Display arr1 len=%d cap=%d value=%v \n", len(cars), cap(cars), cars)

	coords := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := range coords {
		for j := range coords[i] {
			fmt.Printf("coords[%d][%d] = %d\n", i, j, coords[i][j])
		}
	}

	//storing values in array using index
	arr := [...]int{2: 10, 0: 20, 3: 30, 1: 40}
	fmt.Printf("Display arr2 len=%d cap=%d value=%v \n", len(arr), cap(arr), arr)
	//Display values in array using index
	for index, value := range arr {
		fmt.Printf("arr[%d] = %d\n", index, value)
	}
	//Slicing an array
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	weekends := days[5:]
	fmt.Printf("Display weekends %v", weekends)
}
