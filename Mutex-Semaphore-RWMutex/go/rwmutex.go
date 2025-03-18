package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Reader(wg *sync.WaitGroup, rwm *sync.RWMutex, data *int, id int) {
	defer wg.Done()
	fmt.Printf("Reader: %d Waiting\n", id)
	rwm.RLock()
	defer rwm.RUnlock()
	fmt.Printf("Reader: %d, Value: %d\n", id, *data)
	time.Sleep(1 * time.Second)
	fmt.Printf("Reader: %d exiting\n", id)
}

func Writer(wg *sync.WaitGroup, rwm *sync.RWMutex, data *int, dataVal int, id int) {
	defer wg.Done()
	fmt.Printf("--Writer: %d Waiting--\n", id)
	rwm.Lock()
	defer rwm.Unlock()
	*data = dataVal
	fmt.Printf("Current Writer: %d, Value: %d\n", id, dataVal)
	time.Sleep(2 * time.Second)
	fmt.Printf("Writer: %d exiting\n", id)
}

func RunRWMutex() {
	var wg sync.WaitGroup
	var mu sync.RWMutex

	Threads := 15
	data := rand.Intn(7)

	for i := 0; i < Threads; i++ {
		if i > 0 && i%3 == 0 {
			wg.Add(1)
			go Writer(&wg, &mu, &data, rand.Intn(10), i)
		} else {
			wg.Add(1)
			go Reader(&wg, &mu, &data, i)
		}
	}

	wg.Wait()
	fmt.Printf("Final data: %d", data)

}
