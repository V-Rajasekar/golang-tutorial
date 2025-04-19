package main

import (
	"fmt"

	"github.com/V-Rajasekar/golang-tutorial/cube"
)

func main() {
	fmt.Println("Hello, World!")
	// Creating an instance of the Dims struct from the cube package
	dims := cube.Dims{}
	dims.SetSize(10, 20, 30)                             // Set the dimensions of the cube
	fmt.Println("Volume of the cube:", dims.GetVolume()) // Get the volume of the cube
	fmt.Println("Area of the cube:", dims.GetArea())     // Get the area of the cube
}
