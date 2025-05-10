package main

import (
	"fmt"
	"time"
)

func main() {
	dateFunc()
	timeFunc()
}

func timeFunc() {
	t := time.Now()
	fmt.Println("Current Time:", t)
	fmt.Println("Hour:", t.Hour())
	fmt.Println("Minute:", t.Minute())
	fmt.Println("Second:", t.Second())
	fmt.Println("Nanosecond:", t.Nanosecond())
	fmt.Println("Location:", t.Location())
	 hr, min, sec := t.Clock()
	fmt.Printf("Current Time: %02d:%02d:%02d\n", hr, min, sec)
	fmt.Println("Unix Timestamp:", t.Unix())
	fmt.Println("Unix Nano Timestamp:", t.UnixNano())
	fmt.Println("Weekday:", t.Weekday())
}

func dateFunc() {
	dt := time.Now()
	fmt.Println("Current Date and Time:", dt)
	fmt.Println("Year:", dt.Year())
	fmt.Println("Month:", dt.Month())
	fmt.Println("Day:", dt.Day())
	fmt.Println("Hour:", dt.Hour())
	fmt.Println("Minute:", dt.Minute())
	fmt.Println("Second:", dt.Second())
	fmt.Println("Weekday:", dt.Weekday())
}