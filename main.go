package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()
	respch := make(chan any, 10)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fetchUserLikes(userName, respch)
		wg.Done()
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fetchUserMatch(userName, respch)
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println("all fetching is complete")
	close(respch)
	fmt.Println("closed respch")

	for resp := range respch {

		fmt.Println("resp: ", resp)
	}

	fmt.Println("took: ", time.Since(start))
}

// func printResponse(respch chan any) {
// 	for resp := range respch {
// 		fmt.Println("resp: ", resp)
// 	}
// }

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "BOB"
}

func fetchUserLikes(userName string, respch chan any) {
	time.Sleep(time.Millisecond * 150)

	respch <- 11
}

func fetchUserMatch(userName string, respch chan any) {
	time.Sleep(time.Millisecond * 100)

	respch <- "ANNA"
}
