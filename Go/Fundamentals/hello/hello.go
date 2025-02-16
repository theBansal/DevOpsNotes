package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Yash", "Raj", "Rahul", "Rohit", "Rohan"}

	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	for key, value := range messages {
		fmt.Println(key, value)
		fmt.Printf("\n")
	}

}
