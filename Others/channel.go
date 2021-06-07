package main

import "fmt"

func main() {
	// c := make(chan int) // unbuffered channel, error, because channel not initialized yet when call by :9
	c := make(chan int, 1) // buffered channel, success, because channel initialized

	c <- 42
	// go func() { // place it inside go func, so unbuffered channel value call will be after channel initialized
	// 	c <- 42
	// }()

	v, ok := <-c
	fmt.Println(v, ok)
}
