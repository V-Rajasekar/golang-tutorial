# Pointers

- Go is a call-by-value language, the values passed to functions are copies
- Nonpointer types like primitives, structs, and arrays, this means that the called function cannot modify the original.

## Creating a pointer and fetch the value from the pointer
```go
	x := "hello"
	//& address operation. It precedes the value type and returns address
	pointerX := &x
	fmt.Println("simple value x: ", x)
	fmt.Println("&x Address: ", pointerX)
	//* is the indirect operation, it precedes a variable of pointer type and returns the pointed-to value.
	fmt.Println("*pointerX value: ", *pointerX)

//simple value x:  hello
//&x Address:  0x14000010040
//*pointerX value:  hello

```
