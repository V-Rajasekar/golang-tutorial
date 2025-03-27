package main

import (
	"errors"
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	r, rem, _ := divAndRemainder(5, 2)
	s := fmt.Sprintf("Div 5/2 gives Res:%d, Reminder:%d", r, rem)
	fmt.Println(s)

	fmt.Println("creating a anaymous function")
	f := func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")
	}
	for i := 0; i < 5; i++ {
		f(i)
	}
	//closures
	a := 20
	b := 40
	cf := func() {
		fmt.Println(a)
		fmt.Println(b)
		a = 30
		b := 50 //shadow var
		fmt.Println(b)
	}
	cf()
	fmt.Println(a)
	fmt.Println(b)
	//Passing functions as Parameters
	var persons []Person = createPersonSlice()
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].LastName < persons[j].LastName
	})
	fmt.Println(persons)

	//Go Is Call by value
	p := Person{}
	i := 2
	str := "Hello"
	modifyFails(i, str, p)
	fmt.Println(i, str, p)

	//modify map and slice
	m := map[int]string{
		1: "Hi",
		2: "Good",
		3: "Morning",
	}
	fmt.Println("Before map modification:", m)
	modMap(m)
	fmt.Println("After:", m)

	mslice := []int{1, 2, 3}
	fmt.Println("Before slice modification:", mslice)
	modSlice(mslice)
	fmt.Println("After:", mslice)

}

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func createPersonSlice() []Person {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)
	return people
}

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
