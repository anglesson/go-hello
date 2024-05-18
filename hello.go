package main

import (
	"fmt"
	"log"
	"net/http"

	"anglesson/greetings"

	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	http.Handle('/', )
	fmt.Println(messages)
	fmt.Println(quote.Go())
}

func Hello(w *http.ResponseWriter, r http.Request) {
	w.
}
