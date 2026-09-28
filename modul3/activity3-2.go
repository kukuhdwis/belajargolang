package main

import (
	"fmt"
	"sync"
)

// Shared variable
var counter int

// Mutex used to synchronize access to 'counter' and prevent race conditions
var mutex sync.Mutex

// safeIncrement increments the shared variable safely using a mutex
func safeIncrement(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		// Acquire lock before entering the critical section
		mutex.Lock()
		counter++ // Safely executed under mutual exclusion
		// Release lock after critical section completes
		mutex.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Running activity3-2: Two goroutines synchronized with sync.Mutex...")

	// Launch two concurrent goroutines
	wg.Add(2)
	go safeIncrement(&wg)
	go safeIncrement(&wg)

	// Wait for both goroutines to complete
	wg.Wait()

	fmt.Printf("Expected Counter: %d\n", 20000)
	fmt.Printf("Actual Counter  : %d\n", counter)

	if counter == 20000 {
		fmt.Println("Success: Race condition eliminated! All 20,000 updates were safely applied.")
	} else {
		fmt.Printf("Unexpected mismatch: Got %d\n", counter)
	}
}
