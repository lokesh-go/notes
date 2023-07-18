package blockingcall

import "fmt"

// Example of main task waits until the asynchronous task is complete.
func WithUnBuffChannel() {
	var num []int

	// Unbuffered channel
	done := make(chan struct{})

	// Excecute goroutine
	go func() {
		num = make([]int, 2)

		// sends signal after completing task inside the go routines
		done <- struct{}{}
	}()

	// Blocking the main goroutine until goroutine will be finish their task
	<-done
	num[0] = 1
	fmt.Println(num[0])
}
