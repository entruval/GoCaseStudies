package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i <= 5; i++ {
			c <- i
		}
		close(c) // if you not close channel it will expect to receive more value, and return error
	}()

	for v := range c { // iterate channel
		fmt.Println(v)
	}
}
