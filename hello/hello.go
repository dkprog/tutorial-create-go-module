package main

import (
	"fmt"
	"log"

	"github.com/dkprog/tutorial-create-go-module/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Daniel", "Howard", "Rearden", "Dagny"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
