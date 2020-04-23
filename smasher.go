package main

import (
	"fmt"
	"net/http"
)

type Smasher struct {
}

func (s Smasher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
