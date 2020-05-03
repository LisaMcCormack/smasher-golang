package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubSmasher struct {
}

func (s *StubSmasher) getBody(url string) (result string) {
	if url == "bob.com" {
		result += "bob"
	}
	if url == "cat.com" {
		result += "cat"
	}
	return result
}


func TestSmasher(t *testing.T) {

	t.Run("it gets a response body from a url and puts it on the server response", func(t *testing.T) {
		server := NewServer(&StubSmasher{})
		request, _ := http.NewRequest(http.MethodGet, "/?urls=bob.com", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		if response.Body.String() != "bob" {
			t.Errorf("expected bob, got %q", response.Body)
		}
	})

	t.Run("can handle 2 urls", func(t *testing.T) {
		server := NewServer(&StubSmasher{})
		request, _ := http.NewRequest(http.MethodGet, "/?urls=bob.com&urls=cat.com", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		if response.Body.String() != "bobcat" {
			t.Errorf("expected bobcat, got %q", response.Body)
		}
	})

}
