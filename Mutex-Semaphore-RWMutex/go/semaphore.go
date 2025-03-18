package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func limitEntry(id int, sem *semaphore.Weighted, wg *sync.WaitGroup) {
	defer wg.Done()
	waitTime := time.Duration(rand.Intn(6)+1) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), waitTime)
	defer cancel()

	fmt.Printf("%d - is waiting on the queue\n", id)
	if err := sem.Acquire(ctx, 1); err != nil {
		fmt.Printf("%d - Worker went back, since he waited too long\n", id)
		return
	}

	fmt.Printf("%d - Worker started\n", id)
	sleepTime := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(sleepTime)
	fmt.Printf("%d - Worker finished task\n", id)
	sem.Release(1)
}

func RunSemaphore() {
	fmt.Printf("Running Semaphore. \n\n")
	maxEntry := 15
	var wg sync.WaitGroup
	var sem = semaphore.NewWeighted(3)

	wg.Add(maxEntry)
	for i := 0; i < maxEntry; i++ {
		go limitEntry(i, sem, &wg)
	}
	wg.Wait()
	fmt.Printf("\nExiting Semaphore. \n\n")
}
