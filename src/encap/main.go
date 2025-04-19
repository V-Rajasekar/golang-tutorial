package main

import (
	"fmt"
	"cube"
)

func main() {
	fmt.Println("Hello, World!")
	cube.SetSize(2, 3, 4)
	fmt.Println("Volume:", cube.GetVolume())
	fmt.Println("Area:", cube.GetArea())
}