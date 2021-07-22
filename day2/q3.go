package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankAccount struct {
	fund int
}

func (c *BankAccount) WithDraw(amount int, mu *sync.Mutex) {
	mu.Lock()
	if c.fund >= amount {
		c.fund -= amount
	} else {
		panic("Not, Sufficient Amount")
	}
	mu.Unlock()
}

func (c *BankAccount) Deposit(amount int, mu *sync.Mutex) {
	mu.Lock()
	c.fund += amount
	mu.Unlock()
}

func main() {
	var mu sync.Mutex
	c := BankAccount{fund: 500}
	for i := 0; i < 500; i++ {
		if rand.Intn(2) == 1 {
			go c.WithDraw(rand.Intn(1000), &mu)
		} else {
			go c.Deposit(rand.Intn(10000), &mu)
		}
	}

	time.Sleep(time.Second * 2)
	fmt.Println(c.fund)

}
