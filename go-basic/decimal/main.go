// hexadecimal-octal-ascii-utf8-unicode-runes/hexa-lower/main.go
package main

import "fmt"

func main() {
    n := 2548
    n2 := 0x9F4
    fmt.Printf("%x\n", n)
    fmt.Printf("%X\n", n2)
    fmt.Printf("%X\n Hex format to Binary : %b\n", n2, n2)

    fmt.Println("Hexa octal to decimal")
    hexa_octal_to_decimal()
    fmt.Println("Decimal to Octal with prefix")
    decimal_to_octal_withPrefix()

    addByte()
}

func hexa_octal_to_decimal() {

    n2 := 0x9F4
    fmt.Printf("Decimal : %d\n", n2)

    // n3 is represented using the octal numeral system
    n3 := 02454
    // alternative : n3 := 0o2454

    // convert in decimal
    fmt.Printf("decimal: %d\n", n3)
}

func decimal_to_octal_withPrefix() {

    // n4 is represented using the decimal numeral system
    n4 := 1324
    // output n4 (decimal) in octal
    fmt.Printf("octal: %o\n", n4)
    // output n4 (decimal) in octal (with a 0o prefix)
    fmt.Printf("octal with prefix : %O\n", n4)
}

func addByte() {
    b := make([]byte, 0)
  //  b = append(b, 256) Numeric overflow as byte range (2^8)-1 is 20.255 
    b = append(b, 10)
    fmt.Println(b)
}