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
```
## Pointers Mutable parameter case
```go
var pointerToY *int // creating a new Pointer type, alternative var pointer = new(int)
fmt.Println("new pointerToY:", pointerToY == nil)
pointerToY = &y
fmt.Println("pointerToY pointing to Y address(&y):", pointerToY)
fmt.Println("pointerToY pointing to value:(*pointerToY)", *pointerToY)
z = 30
pointerToY = &z
fmt.Println("pointerToY pointing to z value:", *pointerToY) // 30
z =40
fmt.Println("pointerToY pointing to z value:", *pointerToY) // 40
//simple value x:  hello
//&x Address:  0x14000010040
//*pointerX value:  hello

```

- Pointers Mutable Param case 2 inside a method

```go
func failedUpdate(px *int) {
    x2 := 20
    px = &x2
}

//Dereferencing puts the new value in the memory location pointed to by both the original and the copy
func update(px *int) {
    *px = 20
}
x = 10
failedUpdate(&x) //after call still x=10
udate(&x) // after call x=20
```

## Pointer param usecase
- use pointer parameters to modify a variable is when the function expects an interface. (e.g) Unmarshal([]byte, v any) //v any is a *pointer

```go
f := struct {
  Name string `json:"name"`
  Age int `json:"age"`
}{}
err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &f)
```
