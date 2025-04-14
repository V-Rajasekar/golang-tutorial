package main

import "fmt"

func main() {
	fmt.Println("Hello")

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	stringLength, _ := fmt.Println(str1, str2, str3) // _ dont want to care about err

	// if err == nil {
	fmt.Println("String length:", stringLength)
	// }

	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber)) // %.2f given floating number with 2 decimal point.

	fmt.Printf("Data types: %T, %T, %T, %T, and %T\n",
		str1, str2, str3, aNumber, isTrue)

	myString := fmt.Sprintf("Data types as var: %T, %T, %T, %T, and %T \n",
		str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString)

	var i int = 42
	j := i     // int
	f := 3.142 // float64
	fc := 3.142 * 3.142
	g := 0.867 + 0.5i // complex128

	fmt.Printf("v is of type  %T %T %T %T\n", j, f, fc, g)

	var complexNum = complex(20.3, 10.2)
	fmt.Printf("complexNum: %v of type %T \n", complexNum, complexNum)

	var x int = 10
	var y float64 = 30.2
	var sum = x + int(y) //explicit conversions
	fmt.Printf("Sum %v", sum)
	var z float64 = y + 5 //literals in code are untypes no type cast required

	x = 20                      //reassigning
	var num1, msg = 20, "hello" //multi
	num2, msg2 := 30, "world"

	fmt.Printf("Message %v %v %v %v %v \n", num1, msg, num2, msg2, z)

	const ten int64 = 10

	const (
		idKey   = 10
		nameKey = "name"
	)

	const result = 20 * 10

	const str4 = "Hello"
	const str5 = "World!"
	const result2 = str4 + str5

	fmt.Println("const", result2)

	constant()

}

func constant() {
	const x = 10

	var y int = x
	var z float64 = x
	var d byte = x

	fmt.Printf("%v %v %v", y, z, d)

}
