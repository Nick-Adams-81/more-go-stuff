package main

import (
	"fmt"

	"example.com/greetings/greetings"
)

func main() {

	message := greetings.Hello("Nick")
	fmt.Println(message)

	addition := greetings.Add(3, 6)
	fmt.Println(addition)

	subtraction := greetings.Subtract(10, 4)
	fmt.Println(subtraction)

}
