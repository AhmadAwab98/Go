package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// function Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// use Lock so only one goroutine at a time can access the map
	c.v[key]++
	c.mu.Unlock()
}

// function Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// use Lock so only one goroutine at a time can access the map
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("myKey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("myKey"))
}

