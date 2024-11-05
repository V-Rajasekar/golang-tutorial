package main

import "fmt"

func main() {
predefinedType()
}

func predefinedType() {
	var flag bool // false
    toggleOn := true
    var (
        x int = 10
		
        name string
    )
	
	//var vs := (skip var)
	y  := 20 //dcl & initalize 
	y = 30 //assign valye

	var a, b = 10, "hello"
	c, d := 10, "hello"
	fmt.Print(a, b, c, d)
	
	fmt.Printf("flag %v, y %v, x: %v, name: %v", flag, y, x, name)
	fmt.Printf("Print its types: flag %T, y %T, x: %T, name: %T", y, toggleOn, x, name)
	
	var complexNum = complex(20.3, 10.2)
	fmt.Println("complexNum", complexNum)
	aNumber := 42
	fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber)) // %.2f given floating number with 2 decimal point.

	//How to create constants 
	const ten int64 = 10

	const (
		idKey   = 10
		nameKey = "name"
	)

	const result = 20 * 10
	//const resultNum = x + y compile error x, y has to be constant
	 

}

 