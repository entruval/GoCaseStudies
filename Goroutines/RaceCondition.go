//	race condition is main func finished before it's goroutine

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU:\t\t", runtime.NumCPU())            // number of physical threads
	fmt.Println("Goroutines:\t", runtime.NumGoroutine()) // number of goroutines processed

	go loop1()
	go loop2()
	go loop3()
	go loop4()
	loop5()
	loop6()

	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
}

func loop1() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop1:", i)
	}
}

func loop2() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop2:", i)
	}
}

func loop3() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop3:", i)
	}
}

func loop4() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop4:", i)
	}
}

func loop5() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop5:", i)
	}
}

func loop6() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop6:", i)
	}
}

// results
// CPU:		 4
// Goroutines:		 1
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
// Goroutines:		 2

// loop 1 not printed
