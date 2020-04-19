package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Smasher struct {}


func (s Smasher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urls := r.URL.Query()["urls"]
	smashedBody := s.SmashBodies(urls)
	fmt.Fprint(w, smashedBody)
}

func (s Smasher) SmashBodies(urls []string) string {
	resp, err := http.Get(urls[0])
	if err != nil {
		fmt.Errorf("error making get request to url %s, %v", urls[0], err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	r := string(body)
	return r
}

func main() {
	server := &Smasher{}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

