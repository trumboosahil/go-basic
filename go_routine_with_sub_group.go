package main

import (
	"fmt"
	"sync"
	"time"
)

// Mock function to simulate a database call
func mockDBCall(id int) string {
	// Simulate a delay
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("Result from DB call %d", id)
}

// Function to run as a goroutine
func fetchData(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup that this goroutine is done
	result := mockDBCall(id)
	fmt.Println(result)
}

func main() {
	var wg sync.WaitGroup
	start := time.Now() // Start time

	// Start multiple goroutines to fetch data
	for i := 1; i <= 1000; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go fetchData(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Printf("Execution time with goroutines: %v\n", time.Since(start))
}
