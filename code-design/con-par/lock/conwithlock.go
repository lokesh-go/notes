package lock

import (
	"fmt"
	"sync"
)

var totalAmount int

func ConWithLock() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increaseAmount(&wg, &m)
	}
	wg.Wait()
	fmt.Println("Total: ", totalAmount)
}

func increaseAmount(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	totalAmount++
	mutex.Unlock()
}
