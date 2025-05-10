package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter Text: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Println("Input:", scanner.Text())
	} else if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("Enter Words:", scanner.Text())
	words := strings.Fields(scanner.Text())
	fmt.Println("Words:", words)
	fmt.Println("Number of words:", len(words))	
}
