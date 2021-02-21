package main

import (
	"fmt"
	"sync"
)

var mx sync.Mutex

func main() {
	counter := 0

	for i := 0; i < 100; i++ {
		wg := sync.WaitGroup{}
		wg.Add(1)
		go increment(&counter, &wg)
		wg.Wait()
	}
}

func increment(counter *int, wg *sync.WaitGroup) {

	defer wg.Done()
	mx.Lock()
	v := *counter
	v++
	*counter = v
	fmt.Println(*counter)
	mx.Unlock()
}
