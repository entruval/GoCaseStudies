package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	goodDay   func()
}

type client interface {
	greeting()
}

func main() {
	henry := user{
		firstName: "John",
		lastName:  "Doe",
		goodDay: func() {
			fmt.Println("Good day")
		},
	}
	henry.greeting()
	henry.goodDay()
}

func (user *user) greeting() {
	fmt.Println("Hello my name is", user.firstName, user.lastName)
}
