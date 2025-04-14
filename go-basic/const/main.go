package main

import (
	"fmt"
)

func main() {
	const (
		red   = iota // 0
		green        // 1	
		blue         // 2		
	)
	fmt.Println(red, green, blue) // 0 1 2
}