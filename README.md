- [golang-tutorial](#golang-tutorial)
  - [Go Predefined types](#go-predefined-types)
  - [var vs :=](#var-vs-)
  - [Slices](#slices)
  - [Strings Runes and Bytes](#strings-runes-and-bytes)
  - [Maps](#maps)
  - [Structs](#structs)
  - [Functions](#functions)
  - [Pointers](#pointers)
  - [Resources](#resources)


# golang-tutorial
Learning GoLang fundamentals with examples

## Go Predefined types
Predeclared types are

- bool
- string
- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // alias for uint8
- rune // alias for int32
  // represents a Unicode code point (e.g) single Unicode characters ('a'), 8-bit octal numbers ('\141'), 8-bit hexadecimal numbers ('\x61'), 16-bit hexadecimal numbers ('\u0061'), or 32-bit Unicode numbers ('\U00000061')
- float32 float64
- complex64 complex128

`Default values: bool - false, int - 0, string - ""`

```go
//Declaring and intializing types and variables
    var flag // false
    toggleOn := true
    var (
        x int = 10
        name string
    )
```

## var vs :=

- `var` is for declaring a variable `var flag bool`
- `:=`  is declaring and initializing a variable `y := 20`.
- `=` for assigning already declared or initalized variable

```go
    var flag bool // false

    y  := 20 //dcl & initalize 
    y = 30 //assign value

    const x int64 = 10 // typed constant
    cont z = 10 // untyped constant
    //use const to create constants
    var (
    idKey   = "id"
    nameKey = "name"
    )

```

## Slices

- slice length `len(slice)`
- slice append `append(slice, value)`
- capacity `cap(slice)`
- Emptying a Slice `clear(slice)`
- Declaring Your Slice

```go
    var slice []int // slice is nil
    var slice = []int{} // slice with zero element
    make([] int, 5, 10)// empty slice with default value
```

- Slicing Slices
  - Index starts with 0. start inclusive, and end exclusive

  ```go
    Slicing slices [a, b, c, d]
    x: [a b c d]
    y := x[:2] [a b]
    z := x[1:] [b c d]
    d := x[1:3] [b c]
    e := x[:] [a b c d]
  ```

-copy a slice `num := copy(dest, src)` num no of copied values
>note you can use copy with arrays, and slice and interchangeably.
- Converting Arrays to Slices.


```go
    xArray := [4]int{5, 6, 7, 8}
    xSlice := xArray[:]
    //convert a subset of an array into a slice:
    x := [4]int{5, 6, 7, 8}
    y := x[:2]
```

    - Taking a slice from an arrray has the same memory-sharing properties as taking a slice from a slice. Any update in original slice, affects the copied slice and vice versa  
- Converting Slices to Arrays
  - size of the array can be smaller, but not higher than slice size

  ```go
  xSlice := []int{1, 2, 3, 4}
  xArray := [4]int(xSlice)
  ```

## Strings Runes and Bytes

- String is composed of a sequence of UTF-8 encoded code points.
- String len computed using `len(s)`
- Extracting a single value or a group of value is possible with index expression

```go 
  var s string = "Hello there"
  var s2 string = s[4:7] // o t
  var s3 string = s[:5]// Hello
  var s4 string = s[6:]// there
```

- Converting string to byte `var bs [] byte= []byte(s)`

- Runes
- A single unicode char is called Runes 'x' like 'y'
-  A single rune or byte can be converted to a string: `var b rune = 'y'`
   `var s2 = string = string(b)`

## Maps

- create nilMap `var nilMap map[String]int` attempting to read returns the zero value for the map's value type. while write cause *panic*
- create empty map `totalWins := map[string]int{}` read and write allowed
- create map with default size:  `make(map[string], int, 10)`
- Comparing Maps `maps.Equal(map1, map2) //prints true irrespective of the element in the order`
- `comma ok Idiom` to check a key present in map

  ```go
    m := map[string]int{
        "hello": 5,
        "world": 0,
    }
    //ok is true/false depending on the key present in the map.
    v, ok := m["hello"]
    fmt.Println(v, ok) // 5 true
  ```

- There is no set in GoLang. you have to create a map with either boolean, or struct value and your element as key.

  ```go
    func mapAsSetWithStruct() {
    fmt.Println("------create set using map with structs ----")
    intSet := map[int]struct{}{}
    vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
      for _, v := range vals {
        intSet[v] = struct{}{}
      }

      if _, ok := intSet[5]; ok {
        fmt.Println("5 is in the set")
      }
    }
  ```

## Structs

- Declaring struct

  ```go 
    type person struct {
      name string
      age  int
      pet  string
    }
  ```

- Define variables of that type: `var fred person`
- Assigning a struct to variable `bob := person{}` values are set to default
- creating struct with value `beth := person{ age:  30, name: "Beth",}`
- Accessing field `bob.name ="Bob"` `fmt.Println(bob.name)`
- Anonymous Structs

- Here the type of variables `person` and `pet` are anonymous structs.
- Common situations: The first is when you translate external data into a struct or a struct into external data (like JSON or Protocol Buffers). This is called unmarshaling and marshaling data, respectively.

  ```go

    var person struct {
      name string
      age  int
      pet  string
    }

    person.name = "bob"
    person.age = 50
    person.pet = "dog"

    pet := struct {
      name string
      kind string
    }{
      name: "Fido",
      kind: "dog",
    }
  ```

  - comparing and converting structs
  - Go does allow you to perform a type conversion from one struct type to another if the fields of both structs have the same names, order, and types, Also no of fields should match.
  - No method to overide unlike equals in java to compare incomparable structs for equality

## Functions

- Declaring and calling functions

- Go allows for multiple return values
- if you weren’t going to read remainder, you would write the assignment as result, _, err := divAndRemainder(5, 2)

```go
   func divAndRemainder(num, denom int) (int, int, error) {
    if denom == 0 {
        return 0, 0, errors.New("cannot divide by zero")
    }
    return num / denom, num % denom, nil
  }

  //calling
  func main() {
    result, remainder, err := divAndRemainder(5, 2)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(result, remainder)
}
```

- Functions are values

- creating a function variable `var myFuncVariable func(string) int`

  ```go
  var myFuncVariable func(string) int
  func f1(a string) int {
    return len(a)
  }
  myFuncVariable = f1
    fmt.Println(myFuncVariable("Hello"))

  ```

- Anonymous Functions

```go
 f := func(j int) {
        fmt.Println("printing", j, "from inside of an anonymous function")
    }
```

- Closures
  - Writing functions inside another function is closures the inner function can access and modify variables declared in the outer function.

 ```go
   func main() {
      a := 20
        f := func() {
            fmt.Println(a)
            a = 30
        }
      f()
      fmt.Println(a)
  }
  //prints 20 30
 ```

- Passing Functions as Parameters
- sort function takes a slice and a function, here creating a anaymous function

```go
  // sort by last name
  sort.Slice(people, func(i, j int) bool {
      return people[i].LastName < people[j].LastName
  })
  fmt.Println(people)
```

- defer

- Use `defer` to release the resources. (e.g) `defer f.close()`. `defer` can be used with a function, method or closure call.
- In a Go function, there can be multiple `defer` functions, they run in
last-in, first-out(LIFO) order;

## Go Is Call by value

- When you supply a variable for a parameter to a function, Go always makes a copy of the variable's value, meaning any changes to primitive type int or string will not affect original.
- In case of map, and slice any delete, update will affect the original. The notable difference b/w map and slice is that you can add a new element in map, but not slice. The reason is map and slice are both implemented with pointers.


```go
  func modifyFails(inum int, s string, p Person) {
    inum = inum * 2
    s = "Good bye"
    p.FirstName = "Bob"
  }

  func modMap(gmap map[int]string) {
    gmap[1] = "Hello,"
    gmap[3] = "Evening"
    delete(gmap, 1)
    gmap[4] = "raj"
  }

  func modSlice(mslice []int) {
    for k, v := range mslice {
      mslice[k] = v * 2
    }
    mslice = append(mslice, 10)
  }
```

## Pointers

GO pointers:

1. & address operation. It precedes the value type and returns address
2. * is the indirect operation, it precedes a variable of pointer type and returns the pointed-to value.
3. fooPointer := &Foo{} use & in case of struct, not applicable to primitive type
4. var pointerToY *int // creating a new Pointer type, alter var pt = new(int)
5. when assigning a value to a pointer variable inside a struct use generic method to take any type and
return the pointerReference see makePointer
6. Go is a call-by-value language the values passed to functions are copies. Following non pointer types value are not affected when its mosified inside a function. For non pointer types(in method param) like primitives, string, struct, and arrays
7. Pointer passing performance: The time to pass a pointer into a function is constant for all data sizes. Passing a value into a function takes longer as the data gets larger  Returning a pointer vs value is if a data 100-byte returning value takes 10ns, pointer 30ns, but when data is large 100MB value takes 1.5 ms, but pointer takes 0.5ms


## Resources

Links for Self-Learning Go

- Practical Go Lessons - https://www.practical-go-lessons.com
- Ultimate Go Programming (Video) - https://www.oreilly.com/library/view/ultimate-go-programming/9780135261651/
- Golang on Pluralsight - https://app.pluralsight.com/paths/skill/go-core-language
- Go By Example - https://gobyexample.com/
- Go course at Codecademy - https://www.codecademy.com/learn/learn-go
- Go Tour - https://tour.golang.org
- Effective Go - https://golang.org/doc/effective_go.html
- Learn Go With Tests - https://quii.gitbook.io/learn-go-with-tests/
