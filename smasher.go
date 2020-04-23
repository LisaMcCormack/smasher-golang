package main

import (
	"fmt"
	"net/http"
)

type Smasher struct{}

func (s Smasher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, s.getBody(r))
}

func (s Smasher) getBody(r *http.Request) string {
	body := "Hello World"
	urls := r.URL.Query()["urls"]
	if len(urls) == 1 {
		body = "Bob"
	}
	return body
}


