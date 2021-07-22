package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, sum *int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(1)))
	*sum += rand.Intn(100)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var wg sync.WaitGroup
	var sum int = 0
	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go worker(&wg, &sum)
	}
	wg.Wait()
	fmt.Println((sum / 200))
}
