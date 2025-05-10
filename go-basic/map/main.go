package main

import (
	"fmt"
)

func main() {
	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"

	fmt.Printf("Display colors len=%d value=%v \n", len(colors), colors)
	fmt.Printf("Display colors[red] = %s\n", colors["red"])
	fmt.Print("Display colors value=")
	for key, value := range colors {
		fmt.Printf("%s=%s ", key, value)
	}

}