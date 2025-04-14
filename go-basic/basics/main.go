package main

import "fmt"

func main() {
	var roomNumber, floorNumber int = 154, 3
	fmt.Println(roomNumber, floorNumber)

	var password = "notSecured"
	fmt.Println(password)

	//short
	aptNumber, doorNumber := 154, 3
	fmt.Println(aptNumber, doorNumber)

	const hotelName string = "Gopher Hotel"

	//no type constants
	const longitude, latitude =  24.806078, -78.243027

	occupancy := 12
	var available uint8 = 255



	fmt.Printf("hotalName %s longitude %v, latitude %v, occupancy %d \n", hotelName, longitude, latitude, occupancy)

	fmt.Printf("occupancy %d of type %T, and available %d of type %T", occupancy,occupancy, available,available)
	 

}
