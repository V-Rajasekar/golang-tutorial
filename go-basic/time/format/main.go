package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	fmt.Println("Current Date and Time:", dt)
	fmt.Println("Formatted Date (YYYY-MM-DD):", dt.Format("2006-01-02"))
	fmt.Println("Formatted Time (HH:MM:SS):", dt.Format("15:04:05"))
	fmt.Println("Formatted Date and Time:", dt.Format("2006-01-02 15:04:05"))
	fmt.Println("US Locale (MM/DD/YYYY):", dt.Format("January 2, 2006"))
	fmt.Println("UK Locale (DD/MM/YYYY):", dt.Format("2 January 2006"))
	fmt.Println("12-Hour Format (hh:mm:ss AM/PM):", dt.Format("03:04:05 PM"))
	fmt.Println("24-Hour Format (HH:mm:ss):", dt.Format("15:04:05"))

}