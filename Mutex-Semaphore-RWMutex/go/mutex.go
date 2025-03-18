package main

import (
	"fmt"
	"sync"
)

func incCounter(wg *sync.WaitGroup, counter *int) {
	defer wg.Done()
	*counter = *counter + 1
}

func threadSafeCounterIncrement(wg *sync.WaitGroup, mu *sync.Mutex, counter *int) {
	defer wg.Done()
	mu.Lock()
	*counter = *counter + 1
	mu.Unlock()
}

func RunMutex() {
	fmt.Printf("Running Mutex. \n\n")
	
	var wg sync.WaitGroup
	var mu sync.Mutex

	counter := 0
	threadCount := 2000000

	wg.Add(threadCount)
	for i := 0; i < threadCount; i++ {
		go incCounter(&wg, &counter)
	}

	wg.Wait()

	fmt.Printf("Non-thread safe counter value is: %d\n", counter)

	counter = 0

	wg.Add(threadCount)

	for i := 0; i < threadCount; i++ {
		go threadSafeCounterIncrement(&wg, &mu, &counter)
	}

	wg.Wait()

	fmt.Printf("Thread-safe counter value is: %d\n", counter)

	fmt.Printf("\nExiting Mutex. \n\n")
}
