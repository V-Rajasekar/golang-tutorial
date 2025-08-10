package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\n Line by line Execution")
	count("moose")
	count("sheep")

	fmt.Println("\n Concurrent Execution")
	go count("moose")
	count("sheep")
}

func count(item string) {
	for i:=1; i<= 3; i++ {
		fmt.Printf("%v %v ", i, item)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}