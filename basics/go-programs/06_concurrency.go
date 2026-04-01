package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	GO PROGRAM 6: CONCURRENCY (GOROUTINES, SYNC & CHANNELS)
	Concepts:
	- 'go' keyword starts a light-weight "goroutine" (thread).
	- 'sync.WaitGroup' ensures the main thread waits for goroutines to finish.
	- 'Channels' are used for safe communication between concurrent tasks.
*/

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this task is finished at the end
	fmt.Printf("Worker %d: Starting work...\n", id)
	time.Sleep(time.Duration(id) * 500 * time.Millisecond) // Simulate work
	fmt.Printf("Worker %d: Finished work.\n", id)
}

func main() {
	var wg sync.WaitGroup
	
	fmt.Println("Concurrency (Running 3 workers concurrently):")
	
	// Launch 3 concurrent goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Tell WaitGroup to wait for one more goroutine
		go work(i, &wg)
	}
	
	// Wait here until all wg.Done() have been called
	wg.Wait()
	fmt.Println("All workers finished. Returning to main.")
	
	// CHANNEL EXAMPLE
	messages := make(chan string) // Create a new channel
	
	go func() {
		messages <- "Greetings from a Channel!" // Send data into the channel
	}()
	
	msg := <-messages // Receive data from the channel
	fmt.Println("\nReceived via Channel:", msg)
}
