package main

import (
	"fmt"

	"example.com/greetings/greetings"
)

func main() {
	message := greetings.Hello("Nick")
	fmt.Println(message)

	addition := greetings.Add(3, 7)
	fmt.Println(addition)

}
