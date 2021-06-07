package main

import "fmt"

func main() {
	c := make(chan int)

	go send_value_to_channel(c)
	receive_value_from_channel(c)
}

func send_value_to_channel(c chan<- int) {
	c <- 42
}

func receive_value_from_channel(c <-chan int) {
	fmt.Println(<-c)
}
