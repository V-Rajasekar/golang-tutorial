package main

import (
	"fmt"
	"time"
)


// main demonstrates working with Go's time package, including creating a specific date,
// adding duration, formatting, and parsing a date string.
func main() {
	// Create a specific date and time
	dt := time.Date(2025, time.June, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println("Date and Time:", dt)

	//ADD 10 hours to the date
	// and print the updated date and time
	dt = dt.Add(10 * time.Hour)
	fmt.Println("After Adding 10 Hours:", dt)
	layout := "2006-Jan-02 03:04:05PM"
	// Format the date and time using a custom layout
	strDt := dt.Format(layout)
	fmt.Println("Formatted Date and Time:", strDt)

	// Parse a date string back to time.Time
	var err error
	dt, err = time.Parse(layout, strDt)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	} else {
		fmt.Println("Parsed Date and Time:", dt)
	}

	//Recognize Zones
	zone := "All"
	loc,_ := time.LoadLocation(zone)
	fmt.Println("Location:", loc)

	// Get the current time in the specified location
	loc, _ = time.LoadLocation("America/New_York")
	currentTime := time.Now().In(loc)
	fmt.Println("Current Time in New York:", currentTime)
}
