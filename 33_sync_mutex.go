package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	store map[string]int
}

func (sf *SafeCounter) Inc(key string) {
	sf.mu.Lock();
	sf.store[key]++;
	sf.mu.Unlock();
}

func (sf *SafeCounter) Value(key string) int {
	sf.mu.Lock();
	defer sf.mu.Unlock();
	return sf.store[key];
}

func syncMutex() {
	c := SafeCounter{store:make(map[string]int)};

	for i := 0; i < 100; i++ {
		go c.Inc("someKey")
	}

	time.Sleep(time.Millisecond);

	fmt.Println(c.Value("someKey"))
}