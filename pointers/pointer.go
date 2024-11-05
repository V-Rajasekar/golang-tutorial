package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
*
GO pointers:
1. & address operation. It precedes the value type and returns address
2. * is the indirect operation, it precedes a variable of pointer type and returns the pointed-to value.
3. fooPointer := &Foo{} use & in case of struct, not applicable to primitive type
4. var pointerToY *int // creating a new Pointer type, alter var pt = new(int)
5. when assigning a value to a pointer variable inside a struct use generic method to take any type and
return the pointerReference see makePointer
 6. Go is a call-by-value language the values passed to functions are copies. For non pointer
 types(in method param) like primitives, struct, and arrays the func can't modify the original.
 7. Pointer passing performance: The time to pass a pointer into a function is constant for all data sizes. Passing a value into a
 function takes longer as the data gets larger
 Returning a pointer vs value is if a data 100-byte returning value takes 10ns, pointer 30ns, but when data is large 100MB value takes 1.5 ms
 , but pointer takes 0.5ms
*/

type Foo struct {
 Field1 string `json:"field_1,omitempty"`
 Field2 int
}

type IntTree struct {
 val int
 left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
 if it == nil {
 return &IntTree{val: val}
 }
 if val < it.val {
 it.left = it.left.Insert(val)
 } else if val > it.val {
 it.right = it.right.Insert(val)
 }
 return it
}

func (it *IntTree) Contains(val int) bool {
 switch {
 case it == nil:
 return false
 case val < it.val:
 return it.left.Contains(val)
 case val > it.val:
 return it.right.Contains(val)
 default:
 return true
 }
}



func main() {

 x := "hello"
 //& address operation. It precedes the value type and returns address
 pointerX := &x
 fmt.Println("simple value x: ", x)
 fmt.Println("&x Address: ", pointerX)
 //* is the indirect operation, it precedes a variable of pointer type and returns the pointed-to value.
 fmt.Println("*pointerX value: ", *pointerX)

 //Pointer type is a type that represents a pointer
 y := 10
 z := 20
 var pointerToY *int // creating a new Pointer type, alternative var pointer = new(int)
 fmt.Println("new pointerToY:", pointerToY == nil)
 fmt.Println("new pointerToY value:", pointerToY)
 pointerToY = &y
 fmt.Println("pointerToY pointing to Y address(&y):", pointerToY)
 fmt.Println("pointerToY pointing to value:(*pointerToY)", *pointerToY)
 pointerToY = &z
 fmt.Println("pointerToY pointing to z address:", pointerToY)
 fmt.Println("pointerToY pointing to z value:", *pointerToY)

 //creating pointer reference using & for structs

 fooPointer := &Foo{}
 fmt.Println("fooPointer: ", fooPointer)
 y = 20
 failedUpdate(&y)
 fmt.Println("Y original 20, and new:", y)
 update(&y)
 fmt.Println("Y original 20, and new:", y)

 f := struct {
 Name string `json:"name,omitempty"`
 Age int `json:"age"`
 }{}
 //The Unmarshal function populates a variable from a slice of bytes containing
 //JSON. It is declared to take a slice of bytes and an any parameter.
 //The value passed in for the any parameter must be a pointer.
 //If it is not, an error is returned.
 fmt.Println("Printing ", f)
 err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &f)
 if err != nil {
 print(err)
 }
 fmt.Println("unmarshal Json to Struct address", f)

 p := structPointer()
 fmt.Println("structWith a pointer:\ntype person struct {\n\tFirstName string\n\tMiddleName *string\n\tLastName string\n}", p)

 fmt.Println("Returning middleName value using (*p.MiddleName):", *p.MiddleName)

 PointerReceiver()

 //Pointer Receivers and Value Receivers
 var c Counter
 doUpdateWrong(c) // Taking a value c
 fmt.Println("in main:", c.String())
 doUpdateRight(&c) // Taking a c pointer
 fmt.Println("in main:", c.String())

 // Code your method for nilInstances
 var it *IntTree
 it = it.Insert(5)
 it = it.Insert(10)
 it = it.Insert(3)
 fmt.Printf("Does 2 exists in %v: %v \n",
 *it, it.Contains(2))

 fmt.Printf("Does 3 exists in %v: %v \n",
 *it, it.Contains(3))
}

type person struct {
 FirstName string
 MiddleName *string
 LastName string
}

func structPointer() person {
 p := person{
 FirstName: "Pat",
 //MiddleName: "Perry", // This line won't compile
 MiddleName: makePointer("Perry"),
 LastName: "Peterson",
 }
 return p
}

// Takes any type and return its pointer to that Type
func makePointer[T any](t T) *T {
 return &t
}

func failedUpdate(px *int) {
 newValue := 10
 px = &newValue
}

func update(px *int) {
 *px = 50
}

// Do's and Don'ts
// Don't use pointer in the method param, only exception when an interface is looking for a pointer for an example
// json.unmarshall([]byte(), v any)
func MakeFooDont(f *Foo) error {
 f.Field1 = "val"
 f.Field2 = 20
 return nil
}

func MakeFoo() (Foo, error) {
 f := Foo{
 Field1: "val",
 Field2: 20,
 }
 return f, nil
}

type Counter struct {
 total int
 lastUpdated time.Time
}

// when you want to update multiple fields of a struct then use setter, otherwise use member
func (c *Counter) Increment() {
 c.total++
 c.lastUpdated = time.Now()
}

func (c Counter) String() string {
 return fmt.Sprintf("total: %d, last updated: %v",
 c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
 c.Increment()
 fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
 c.Increment()
 fmt.Println("in doUpdateRight:", c.String())
}

func PointerReceiver() {
 var c Counter //value type
 fmt.Println(c.String())
 c.Increment() //Go converts c to (&c) pointer type
 fmt.Println(c.String()) // c is converted to (*c).string()
}


