package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const maxWorkers int = 10

type Argument struct {
	Yerda string
}

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	argument := make(chan Argument, 10)

	for i := 0; i < max; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, ctx context.Context) {
			defer wg.Done()
			process(ctx, argument)
		}(&wg, ctx)
	}
	wg.Wait()
}

func process(ctx context.Context, argument chan Argument) {
	for {
		select {
		case <-ctx.Done():
			return
		case arg := <-argument:
			go func(arg Argument) {
				arg = arg
				time.Sleep(time.Second * 2)
				fmt.Println("hello")
			}(arg)
		}
	}
}
