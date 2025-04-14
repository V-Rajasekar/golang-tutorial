package main

import "fmt"

func main() {
	var nights int = 3
	var rate float32 = 145.90
    fmt.Println(computePrice(rate, nights))
	passReference()

}

// compute the price with a 200% margin
func computePrice(rate float32, nights int) (price float32) {
    p := rate * float32(nights)
    p = p * 2
    return
}

func occupancyLevel(occupancyRate float32) string{
	switch  {
	case occupancyRate > 70: return "High"
	case occupancyRate > 20: return "Medium"
	default: return "Low"
	}
}

func passReference() {
	// Pass by reference
	var a, b int = 1, 2
	fmt.Println("Before swap:", a, b)
	swap(&a, &b)
	fmt.Println("After swap:", a, b)
}

func swap(x, y *int) {
	// Swap the values of x and y
	temp := *x
	*x = *y
	*y = temp
}