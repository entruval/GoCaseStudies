package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	v, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(v)
}
