package main

import "fmt"


func main() {
	b := "hello"
	//When you iterate over a string, you will iterate over runes.
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	//A rune represents the Unicode code point of a character.
	var aRune rune = 'A'
	fmt.Println("aRune of A is: ", aRune)
	fmt.Printf("Hex of rune A: %X\n", aRune)
	fmt.Printf("Unicode Code point of %c %U\n", aRune, aRune)
}