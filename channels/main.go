package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//create slice of strings
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	//iterate through links and run the check link function for each link
	for _, link := range links {
		//go creates routine for each check link function so it runs all at the same time
		go checkLink(link, c)
	}
	//i=0, i is less than the length of links print the chan string, add 1 to i when a request is completed
	// for i := 0; i < len(links); i++ {
	// 	go checkLink(<-c)
	// }

	//wait for the channel to return a value then assign it to l
	for l := range c {
		//then run the body of the for loop spawning a go rutine that calls check link, passing the link and then passing the channel
		//go checkLink(l, c)

		//function literal that has access to the  value in memory
		go func(link string) {
			//pause the current loop for  5 seconds
			time.Sleep(5 * time.Second)
			checkLink(link, c)

			//receive l (the new value) and copy to memory
		}(l)
	}
}

//function to check the links
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
