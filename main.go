package main

import (
	"fmt"
	"hello/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	var name string
	
	fmt.Scanln(&name)
	message, err := greetings.Hello(name)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(message)
}
