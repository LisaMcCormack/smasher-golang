package main

import (
	"fmt"
	"net/http"
)

type Smasher struct{}

func (s Smasher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urls := r.URL.Query()["urls"]
	if urls[0] == "bob" {
		fmt.Fprint(w, "Bob")
	} else {
		fmt.Fprint(w, "Hello World")
	}
}
