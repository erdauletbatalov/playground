package main

import (
	"fmt"
	"runtime"
	"time"
)

const max int = 10

func main() {
	ticker := time.NewTicker(time.Second * 4)

	// канал для ограничения кол-ва горутин
	semaphore := make(chan struct{}, max)

	for {
		select {
		case <-ticker.C:
			process(semaphore)
		}
	}
}

func process(semaphore chan struct{}) {
	for i := 0; ; i++ {
		fmt.Println("NumGoroutine:", runtime.NumGoroutine())
		semaphore <- struct{}{}
		go func(i int) {
			defer func() {
				<-semaphore
			}()

			time.Sleep(1 * time.Second)
			fmt.Println(i)
		}(i)
	}
}
