package main

import (
	"log"
	"net/http"
)



func main() {
	server := NewServer(&Smoosher{})

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

