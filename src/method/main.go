package main

import (
	"fmt"
)

type dimension struct {
	length int	
	width  int
	height int
}
type car struct {
	brand string
	color string
	body  string
	dimension dimension // embedding struct
	
}

//Attaching a method to the car struct
func (c car) display() {
	fmt.Printf("Car: %s, color: %s, Body: %s Height: %d \n", c.brand, c.color, c.body, c.dimension.height)
}
func (c car) changeColor(newColor string) {
	c.color = newColor
	fmt.Printf("Car color changed to: %s\n", c.color)
}

func main() {
	porsche := car{brand: "porsche", color: "red", body: "sedan"}
	tDimension := dimension{length: 4, width: 2, height: 1}
	tataPunch := car{"tata", "blue", "hatchback", tDimension}
	porsche.display()
	tataPunch.display()
}