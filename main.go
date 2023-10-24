package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	type value struct {
		sync.Mutex
		id     string
		locked bool
	}

	lock := func(v *value) {
		v.Lock()
		v.locked = true
	}
	unlock := func(v *value) {
		v.Unlock()
		v.locked = false
	}
	move := func(wg *sync.WaitGroup, id string, v1, v2 *value) {
		defer wg.Done()
		for i := 0; ; i++ {
			if i >= 3 {
				fmt.Println("canceling goroutine...")
				return
			}

			fmt.Printf("%v: Goroutine is locking\n", v1.id)
			lock(v1) // <1>

			time.Sleep(2 * time.Second)

			if v2.locked { // <2>
				fmt.Printf("%v: Goroutine is , blocked by %v\n", v1.id, v2.id)
				unlock(v1) // <3>
				continue
			}
		}
	}
	a, b, c := value{id: "Process1"}, value{id: "Process2"}, value{id: "Process3"}
	var wg sync.WaitGroup
	wg.Add(2)
	go move(&wg, "first", &a, &b)
	go move(&wg, "second", &b, &c)
	wg.Wait()
}
