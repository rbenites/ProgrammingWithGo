package main

import (
	"fmt"
	"net/http"
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
	//i=0, i is less than the length of links print the chan string, add 1 to i
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

//function to check the links
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up")
	c <- "Its Up"
}
