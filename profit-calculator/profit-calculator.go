package main

import "fmt"

func main() {

	var revenue, expense, taxRate float64
	fmt.Print("Enter revenue:")
	fmt.Scan(&revenue)
	fmt.Print("Enter expense:")
	fmt.Scan(&expense)
	fmt.Print("Enter taxRate:")
	fmt.Scan(&taxRate)
	
	ebt := revenue - expense
	profit := float64(ebt * (1 - taxRate)/ 100)
	ratio := ebt/profit

	fmt.Println("Earning before tax ", ebt)
	fmt.Println("Profit: ", revenue/ebt)
	
	fmt.Println("Ratio: ", ratio)


}