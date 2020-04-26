package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Server struct {
	smasher Smasher
}

type Smasher interface {
	getBody(url string)
	smasher() string
}

func NewServer(smasher Smasher) *Server {
	return &Server{smasher}
}

func (d *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urls := r.URL.Query()["urls"]
	for _, url := range urls {
		d.smasher.getBody(url)
	}
	fmt.Fprint(w, d.smasher.smasher())
}

type Smoosher struct {
	store []string
}

func (s *Smoosher) smasher() string {
	string := strings.Join(s.store, "")
	return string
}

func (s *Smoosher) getBody(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error getting body: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading body: %v", err)
	}
	s.store = append(s.store, string(body))
}
