package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	 
  hotelName := "Grand hotel"	
  totalRooms := 5
  firstRoomNo := 110

  rand.NewSource(time.Now().UTC().UnixNano())
  roomsOccupied := rand.Intn(totalRooms)
  roomsAvailable := totalRooms - roomsOccupied

  
  fmt.Println("Welcome to " + hotelName)
  fmt.Printf("Number of rooms: %d \n", totalRooms)
  fmt.Printf("Rooms available: %d \n", roomsAvailable)
  fmt.Println("Rooms:")	
  for i := 0; i < roomsAvailable; i++ {
	fmt.Printf("- %d : %d people/ %d nights \n", firstRoomNo + i, rand.Intn(10), rand.Intn(12))
  }
  var occupancyLevel string
   occupancyRate := (float32(roomsOccupied)/ float32(totalRooms)) * 100
  if occupancyRate > 70 {
	occupancyLevel = "High"
  } else if occupancyRate > 20  {
	occupancyLevel = "Medium"
  } else {
	occupancyLevel = "Low"
  }
  fmt.Println("Occupancy level:", occupancyLevel)
  fmt.Println("Occupancy rate:", occupancyRate)

}