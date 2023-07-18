package main

import (
	"fmt"
	"sync"

	"github.com/lokesh-go/notes/code-design/con-par/lock"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // This method indicates the number of goroutines to wait.
	go task1(&wg)
	wg.Wait() // This method blocks the execution of code until the internal counter becomes 0.

	wg.Add(1) // This method indicates the number of goroutines to wait.
	go task2(&wg)
	wg.Wait() // This method blocks the execution of code until the internal counter becomes 0.

	taks3() // This task will executed by the main goroutine

	multiTaksEx()
	lock.ConWithLock()
}

func task1(wg *sync.WaitGroup) {
	defer wg.Done() // This will reduce the internal counter value by 1
	fmt.Println("This is task 1")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("This is task 2")
}

func taks3() {
	fmt.Println("This is task 3")
}

func multiTaksEx() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("This is work", i)
		}()
		// One by one sequential wait
		// wg.Wait()
	}
	// All executes at same time
	wg.Wait()
}
