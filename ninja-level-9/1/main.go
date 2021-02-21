package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)
	go third(&wg)
	go second(&wg)
	first()

	wg.Wait()
}

func first() {
	fmt.Println("Test")
}

func second(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func third(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second)
	time.Sleep(time.Second)
	fmt.Println("third() Done")
}
