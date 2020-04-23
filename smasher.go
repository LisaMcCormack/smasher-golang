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
	getBody(r *http.Request) string
}




func NewServer(smasher Smasher) *Server {
	return &Server{smasher}
}

func (d *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	    body := d.smasher.getBody(r)
		fmt.Fprint(w, body)
}

type Smoosher struct {}

func (s Smoosher) getBody(r *http.Request) string {
	urls := r.URL.Query()["urls"]
	resp, _ := http.Get(urls[0])
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}


