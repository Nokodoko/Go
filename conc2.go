package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go boss(c1)
	go minions(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("finished")
}

func boss(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func minions(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		func(v2 int) {
			c2 <- work(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func work(i int) int {
	return i * rand.Intn(500)
}
