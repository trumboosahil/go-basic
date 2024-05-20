package main

import (
	"fmt"
	"time"
)

// Mock function to simulate a database call
func mockDBCall(id int) string {
	// Simulate a delay
	time.Sleep(time.Duration(id) * time.Second)
	return fmt.Sprintf("Result from DB call %d", id)
}

func main() {
	start := time.Now() // Start time

	// Fetch data sequentially
	for i := 1; i <= 3; i++ {
		result := mockDBCall(i)
		fmt.Println(result)
	}

	fmt.Printf("Execution time without goroutines: %v\n", time.Since(start))
}
