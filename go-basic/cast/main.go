package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
	// Convert string to int
	r, err := strconv.Atoi("123") // Convert string to int
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Converted value: %v type: %T \n", r, r)
	}
	//Convert numeric to string
	strconv.Itoa(123) // Convert int to string
	s := strconv.FormatBool(true)
	fmt.Println(s)
	s = strconv.FormatFloat(3.1415, 'f', -1, 64)
	fmt.Printf("Value: %v type: %T \n", s, s)
	s = strconv.FormatFloat(3.1415, 'f', 2, 64)
	fmt.Println(s)
	s = strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(s)
	s = strconv.FormatInt(-42, 16)
	fmt.Println(s)
	s = strconv.FormatUint(42, 16)
	fmt.Println(s)

	// Convert string to float
	f, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Converted value: %v type: %T \n", f, f)
	}

	i := float64(3) // Convert int to float
	fmt.Printf("Value: %v type: %T \n", i, i)
	//int64("34") // Convert float to int
}
