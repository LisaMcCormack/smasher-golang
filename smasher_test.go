package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubSmasher struct{}

func (s StubSmasher) getBody(_ *http.Request) string {
	return "bob"
}

func TestSmasher(t *testing.T) {

	t.Run("it gets a response body from a url and puts it on the server response", func(t *testing.T) {
		server := NewServer(StubSmasher{})
		request, _ := http.NewRequest(http.MethodGet, "/?urls=bob", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		if response.Body.String() != "bob" {
			t.Errorf("expected Bob, got %q", response.Body)
		}
	})
}
