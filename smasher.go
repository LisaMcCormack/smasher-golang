package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Server struct {
	smasher Smasher
}

type Smasher interface {
	getBody(url string) string
}

func NewServer(smasher Smasher) *Server {
	return &Server{smasher}
}

func (d *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result := ""
	urls := r.URL.Query()["urls"]
	for _, url := range urls {
		body := d.smasher.getBody(url)
		result += body
	}
	fmt.Fprint(w, result)
}

type Smoosher struct {}



func (s *Smoosher) getBody(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error getting body: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading body: %v", err)
	}
	return string(body)
}
