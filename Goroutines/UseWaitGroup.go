//	race condition is main func finished before it's goroutine

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	runtime.GOMAXPROCS(4)                                // set number of physical threads to be used
	fmt.Println("CPU:\t\t", runtime.NumCPU())            // number of physical threads
	fmt.Println("Goroutines:\t", runtime.NumGoroutine()) // number of goroutines processed

	// .Add(repetition_count) repetition_count is repetition countdown, substract by 1 each wg.Done()
	// if you set more than needed it will return "fatal error all goroutines are asleep - deadlock"
	// .Done() on others goroutines will be counted too
	wg.Add(10)

	go loop1()
	time.Sleep(1000 * time.Nanosecond)

	wg.Wait()
	loop2()

	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
}

func loop1() {
	incrementer := 0

	for i := 0; i < 10; i++ {
		runtime.Gosched() // trigger the scheduler to run any one of the "ready" goroutines in an intentionally unspecified selection order.

		// mutex used to prevent overlapping update by goroutines
		// it will make goroutines wait one another (synchronous)
		mutex.Lock()
		temp := incrementer
		temp++
		mutex.Unlock()

		fmt.Println("loop1:", i)
		wg.Done()
	}
}

func loop2() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop2:", i)
	}
}

// result
// CPU:		 4
// Goroutines:	 1
// loop2: 0
// loop2: 1
// loop2: 2
// loop2: 3
// loop2: 4
// loop2: 5
// loop2: 6
// loop2: 7
// loop2: 8
// loop2: 9
// CPU:		 4
// Goroutines:	 2
// loop1: 0
// loop1: 1
// loop1: 2
// loop1: 3
// loop1: 4
// loop1: 5
// loop1: 6
// loop1: 7
// loop1: 8
// loop1: 9

// loop 1 printed but after main func finished
