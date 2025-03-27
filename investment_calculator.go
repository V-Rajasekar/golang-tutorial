package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Investment Return calculator")
	var investmentAmount, expectedInterest, period float64
	inflationRate:= 6.0
	fmt.Print("Enter investment amount:")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter expected interest")
	fmt.Scan(&expectedInterest)
	fmt.Print("Enter the periods in year, months")
	fmt.Scan(&period)
	fmt.Print("Enter the inflationRatio default is 6%: ")
	fmt.Scan(&inflationRate)
	
	
	var futureValue = futureValue(investmentAmount, period, expectedInterest)

	futureRealValue := futureValue/ math.Pow(1 + inflationRate/100, float64(period))

	fmt.Println("futureValue: ", futureValue)
	fmt.Println("futureRealValue: ", futureRealValue)
}

func futureValue(investAmt float64, years float64, expectedInterest float64) float64 {

	return investAmt * math.Pow(1+expectedInterest/100, years)
}
