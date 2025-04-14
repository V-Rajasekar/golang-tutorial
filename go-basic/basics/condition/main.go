// control-statements/random-seeded/main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
* Sample on random, if else and switch statement, for condition
 */
func main() {
	hotelName := "The Gopher Hotel"
	fmt.Println("Hotel " + hotelName)
	var roomsAvailable int
	//rand requires a seed to generate random nos correctly rand.NewSource
	rand.NewSource(time.Now().UTC().UnixNano())
	var rooms, roomsOccupied int = 100, rand.Intn(100)
	roomsAvailable = rooms - roomsOccupied
	fmt.Println(roomsAvailable, "rooms available")

	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)

	fmt.Println("John is", ageJohn, "years old")
	fmt.Println("Paul is", agePaul, "years old")
	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else if agePaul == ageJohn {
		fmt.Println("Paul and John have the same age")
	} else {
		fmt.Println("Paul is younger than John")
	}

	switchCond(agePaul, ageJohn)

	loopSeveralTime()
	loopWhileTrue()
	breakOutOfLoops()
	goToLabels()

}

func switchCond(agePaul, ageJohn int) {
	fmt.Print("Switch condition")
	// switch expression
	switch {
	case agePaul > ageJohn:
		fmt.Println("agePaul > ageJohn")
	case agePaul == ageJohn:
		fmt.Println("agePaul == ageJohn")
	case agePaul < ageJohn:
		fmt.Println("agePaul < ageJohn")
	}

	// switch statement; expression
	switch ageSum := ageJohn + agePaul; ageSum {
	case 10:
		fmt.Println("ageJohn + agePaul = 10")
	case 20, 30, 40: //*\label{switchMulti}
		fmt.Println("ageJohn + agePaul = 20 or 30 or 40")
	case 2 * agePaul:
		fmt.Println("ageJohn + agePaul = 2 times agePaul")
	default:
		fmt.Println("Default sum: ", ageSum)
	}
}

func loopWhileTrue() {
	fmt.Println("Loop while true")
	// Infinite loop
	i := 5
	for {
		if (i < 1) {
			fmt.Println("\t\t\t\t Lift Off!")
			break // Break out of the infinite loop
		}
			fmt.Println("\t\t Countdown", i)
			i--
	}
}

func loopSeveralTime() {
	fmt.Print("Loop several times")
	for i := 0; i < 3; i++ { //*\label{forWithSingle1}
		fmt.Println("Outer looping", i)
		for i := 0; i < 2; i++ { //*\label{forWithSingle1}
			fmt.Println("\t Inner looping", i)
		}
	}	
}

func breakOutOfLoops() {
	fmt.Println("Break out of loops")
	for i := 0; i <=3 ; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Printf("\t\t Breaks when i=%d and j=%d\n", i, j)
			
				break // Break out of the inner loop
			}
			if i == 3 && j == 2 {
				fmt.Println("Continue when i=3 and j=2")
				continue 
			}
			fmt.Println("Running i:", i, "j:", j)
		}
	}
}


func goToLabels() {
	fmt.Println("Go to labels")
	for i := 0; i < 5; i++ {
		fmt.Println("Running i=", i)
		if i==2 {
		  goto end	
		}
	}
	end:
}
