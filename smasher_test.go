package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type StubSmasher struct {
	store []string
}

func (s *StubSmasher) getBody(url string) {
	if url == "bob.com" {
		s.store = append(s.store, "bob")
	}
	if url == "cat.com" {
		s.store = append(s.store, "cat")
	}
}

func (s *StubSmasher) smasher() string {
	return strings.Join(s.store, "")
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
