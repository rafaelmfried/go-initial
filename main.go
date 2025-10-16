package main

import (
	"fmt"

	"hello/greetings"
)

func main() {
	var name string
	
	fmt.Scanln(&name)
	message := greetings.Hello(name)
	fmt.Println(message)
}
