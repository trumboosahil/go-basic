package main

import (
	"fmt"
	"time"
)

func printNumbers(prefix string, n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(prefix, i)
		time.Sleep(500 * time.Millisecond) // Sleep for 500 milliseconds
	}
}

func main() {
	// Start two goroutines
	go printNumbers("Goroutine 1:", 5)
	go printNumbers("Goroutine 2:", 5)

	// Wait for the goroutines to finish
	// This is a simple way to wait; in a real program,
	// you might use sync.WaitGroup or channels.
	time.Sleep(3 * time.Second)

	fmt.Println("Main function finished")
}
