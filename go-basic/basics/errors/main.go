package main

import (
	"fmt"
)

func main() {
	for i := 2; i >= -2; i-- {
		num, err := isPos(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Positive number:", num)
		}
	}
}
