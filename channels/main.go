package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://comcast.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkStatus(link, c)
	}

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkStatus(<-c, c)
	// }

	//cleaner code writing
	// for l := range c {
	// 	// time.Sleep(time.Second * 5) // not efficient
	// 	go checkStatus(l, c)
	// }

	//loop variable l captured by func literal
	// for l := range c {
	// 	go func() {
	// 		time.Sleep(time.Second)
	// 		checkStatus(l, c)
	// 	}()
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkStatus(link, c)
		}(l)
	}
}

func checkStatus(link string, c chan string) {
	// time.Sleep(time.Second * 5) // not efficient
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "not working.")
		c <- link
		return
	}

	fmt.Println(link, "working")
	c <- link
}
