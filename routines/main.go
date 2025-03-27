package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"https://google.com", "http://facebook.com",
		"http://golang.org", "http://amazon.com"}

	c := make(chan string)
	for _, link := range links {
		//go create a new thread for go Routin. In this case 4 threads, go routine are created.
		go checkLink(link, c)
	}
	//This passes the mainRoutine for 5 seconds eventually all the child routine also has to pass 5 secs to resolve this
	//see below solution
	/*for l := range c {
	go checkLink(l, c)
	}*/

	//This is called function literals
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	//time.Sleep(5 * time.Second)
	_, err := http.Get(link) //Blocking call!
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, " is up")
	c <- link
}