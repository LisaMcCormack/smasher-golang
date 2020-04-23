package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Smasher struct {
	bodies []string
}


func (s Smasher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urls := r.URL.Query()["urls"]
	smashedBody := s.SmashBodies(urls)
	fmt.Fprint(w, smashedBody)
}

func (s Smasher) SmashBodies(urls []string) string {
	for _, url := range urls {
		body := s.getBodies(url)
		s.bodies = append(s.bodies, body)
	}
	return strings.Join(s.bodies, "")
}

func (s Smasher) getBodies(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("error making get request to url %s, %v", url, err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func main() {
	server := &Smasher{}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

