package main

import (
	"errors"
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	mobile *int
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
	fmt.Println("Closure execution started")
	a := 20
	b := 40
	cf := func() {
		fmt.Printf("Accessing outer variable 'a' with value: %d\n", a) // Accessing outer variable 'a'
		fmt.Printf("Accessing outer variable 'b' with value: %d\n", b) // Accessing outer variable 'b'
		a = 30 // Modifying outer variable 'a'
		fmt.Printf("Modified outer variable 'a' to value: %d\n", a)
		b := 50 // Shadowing outer variable 'b' with a new local variable
		fmt.Printf("Shadowed local variable 'b' with value: %d\n", b)
	}
	cf()
	fmt.Println("Value of 'a' after closure execution:", a)
	fmt.Println("Value of 'b' after closure execution:", b)
	fmt.Println("Closure execution completed")
	
	//Passing functions as Parameters
	var persons []Person = createPersonSlice()
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].LastName < persons[j].LastName
	})
	fmt.Println("After Sorting by LastName", persons)

	//Go Is Call by value
	p := Person{}
	i := 2
	str := "Hello"
	fmt.Printf("Before calling modifyFails: i = %d, str = %q, p = %+v\n", i, str, p)
	modifyFails(i, str, p)
	fmt.Printf("After calling modifyFails: i = %d, str = %q, p = %+v\n", i, str, p)

	fmt.Printf("Before calling modifyWorks: i = %d, str = %q, p = %+v\n", i, str, p)
	modifyWorks(i, str, &p)
	fmt.Printf("After calling modifyWorks: i = %d, str = %q, p = %+v, mobile = %d\n", i, str, p, *p.mobile)

	//modify map and slice
	fmt.Println("Map and Slice modification")
	m := map[int]string{
		1: "Hi",
		2: "Good",
		3: "Morning",
	}
	fmt.Println("Before map modification:", m)
	modMap(m)
	fmt.Println("After map modification:", m)

	mslice := []int{1, 2, 3}
	fmt.Println("Before slice modification:", mslice)
	modSlice(mslice)
	fmt.Println("Before slice modification:", mslice)

}

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func createPersonSlice() []Person {
	defaultMobile := 0
	people := []Person{
		{"Pat", "Patterson", nil, 37},
		{"Tracy", "Bobdaughter", nil, 23},
		{"Fred", "Fredson", &defaultMobile, 18},
	}
	fmt.Println(people)
	return people
}

func modifyFails(inum int, s string, p Person) {
	inum = inum * 2
	updateMobile :=  48639822
	s = "Good bye"
	p.FirstName = "Bob" //Use pointer receiver to modify the value
	p.LastName = "Bobson"
	p.mobile = &updateMobile
	fmt.Printf("Inside modifyFails: inum = %d, s = %q, p = %+v\n", inum, s, p)
	// The changes to inum, s, and p are not reflected outside this function
}

func modifyWorks(inum int, s string, p *Person) {
    inum = inum * 2
    updateMobile := 48639822
    s = "Good bye"
    p.FirstName = "Bob" // Modifies the original struct
    p.LastName = "Bobson"
    p.mobile = &updateMobile // Updates the original struct's mobile field
    fmt.Printf("Inside modifyWorks: inum = %d, s = %q, p = %+v\n", inum, s, *p)
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
