package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Create a new random number generator
	var num int = rand.Intn(20) + 1
	var guess int = 0
	var flag bool = true

	fmt.Print("Guess a number between 1 and 20: ")

	for flag {
		_, err := fmt.Scan(&guess) // Read user input
		//_, err := fmt.Scanf("Guess:%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
		}
		if guess < num {
			fmt.Println("Too low! Try again.")
		} else if guess > num {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Congratulations! You guessed the number.")
			flag = false // Exit the loop if the guess is correct
		}
	}
}
