package main

import (
	"fmt"
	"sync"
)

// Shared variable accessed concurrently by two goroutines without synchronization
var counter int

// increment adds 1 to the shared variable across multiple iterations
func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		// RACE CONDITION OCCURS HERE:
		// counter++ consists of three distinct machine operations:
		// 1. Read: read current value of counter from memory into register
		// 2. Modify: increment value by 1 in register
		// 3. Write: store incremented value back to memory
		//
		// Because there is no synchronization (e.g. mutex or channel),
		// two goroutines can read the same stale value simultaneously and
		// overwrite each other's updates, resulting in lost increments.
		counter++
	}
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Running activity3-1: Two goroutines with a race condition...")

	// Launch two concurrent goroutines sharing 'counter'
	wg.Add(2)
	go increment(&wg)
	go increment(&wg)

	// Wait for both goroutines to finish
	wg.Wait()

	fmt.Printf("Expected Counter: %d\n", 20000)
	fmt.Printf("Actual Counter  : %d\n", counter)

	if counter != 20000 {
		fmt.Printf("Result: A race condition occurred! Lost %d updates.\n", 20000-counter)
	} else {
		fmt.Println("Result: Ran without lost updates this time. Run with 'go run -race activity3-1.go' to detect data race.")
	}
}
