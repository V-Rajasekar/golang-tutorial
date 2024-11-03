package main

import (
	"fmt"
	"maps"
)

func main() {
	//comparing two arrays using == !=
	arrays()
	slices()
	//index starts with 0 start-inclusive, end-excluded
	fmt.Println("Slicing slices [a, b, c, d]")
	slicingSlices()
	copyArrayToSlice()
	mapSample()
	mapAsSet()
	mapAsSetWithStruct()
	compareStruct()
}

func arrays() {
	var x [3]int

	fmt.Printf("x: %v \n", x)
	x = [3]int{10, 20, 30}

	fmt.Printf("x: %v \n", x)

	var y = [3]int{10, 20, 30}

	fmt.Printf(" x %v= y %v %v \n", x, y, x == y)
	fmt.Println("len(x): ", len(x))
}

/**
*
*
 */
func slices() {
	fmt.Println("----- Slices -------")
	var x []int // creates nil slice
	var emptySlice = []int{}
	fmt.Printf("var emptySlice = []int{} len(emptySlice): %v", len(emptySlice))

	fmt.Printf("Creating nil slice: var x []int %v \n", x)
	fmt.Printf("x is nil: %v \n", x == nil)
	x = append(x, 100, 200, 300)
	fmt.Printf("x = append(x, 100, 200): %v \n", x)
	fmt.Printf("x is nil: %v \n", x == nil)
	y := []int{10, 20, 30}
	fmt.Println("Creating literal \n slice y := []int{10, 20, 30}: ", y)
	p := make([]int, 5)
	q := make([]int, 5, 10)
	fmt.Printf("Create Empty slice \n p: = make([]int, 5 ) with p: %v len(p): %v \n", p, len(p))
	fmt.Println("Slice with len 5 and cap 10: \n q := make([]int, 5, 10 ) q:\n", q)
	var z = []int{10, 20, 30}

	fmt.Printf("y := []int{10, 20, 30}: %v var z = []int{10, 20, 30} %v \n", y, z)

	fmt.Println("len(x): ", len(x))
	fmt.Println("cap(x): ", cap(x))

	slice := []string{"first", "second", "third"}
	fmt.Println(slice, len(slice))
	clear(slice)
	fmt.Println("Empty a slice: \n clear(s) \n Result:", slice)
	fmt.Println("len(slice):", len(slice))

}

func slicingSlices() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y := x[:2]", y)
	fmt.Println("z := x[1:]", z)
	fmt.Println("d := x[1:3]", d)
	fmt.Println("e := x[:]", e)
}

func copyArrayToSlice() {
	fmt.Println("---copyArrayToSlice---")
	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	x[0] = 10
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func mapSample() {
	var nilMap map[int]string

	fmt.Println("Is Map nil:", nilMap == nil)

	emptyMap := map[int]string{}
	fmt.Println("Is Map empty:", len(emptyMap))
	defaultMap := make(map[int]string, 10)
	fmt.Println("Is defaultMap empty:", len(defaultMap))

	teams := map[int]string{
		1: "India",
		2: "Norway",
		3: "UK",
	}
	for k, v := range teams {
		fmt.Printf("%v:%v \n", k, v)
	}
	fmt.Println("Team 1 is: ", teams[1])

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	fmt.Println("map:", m)
	v, ok := m["hello"]
	fmt.Println(v, ok) // 5 true

	v, ok = m["goodbye"]
	fmt.Println(v, ok) // 0 false

	delete(m, "hello") // m is null nothing happens
	fmt.Printf("map:%v, len(m):%v", m, len(m))
	clear(m)
	fmt.Printf("map:%v, len(m):%v", m, len(m))
	dest := map[string]int{
		"hello": 5,
		"world": 0,
	}
	fmt.Println(" maps.Equal(dest, m): ", maps.Equal(dest, m))

}

func mapAsSet() {
	fmt.Println("---- Using map to frame a set ---")

	intSet := map[int]bool{}
	intSlice := []int{5, 6, 9, 10, 100, 200, 9, 6}

	for _, val := range intSlice {
		intSet[val] = true
	}
	fmt.Printf("intSet %v  \n", intSet)
	fmt.Printf("intSlice %v \n", intSlice)
	fmt.Println("len(intSet):", len(intSet))
	fmt.Println("len(intSlice):", len(intSlice))
	fmt.Println(intSet[100])
	fmt.Println(intSet[500])
	if intSet[200] {
		fmt.Println("200 is present")
	}
}

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

	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in the set")
	}
}

func compareStruct() {
	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}

	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g)
}
