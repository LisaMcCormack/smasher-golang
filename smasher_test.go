package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSmasher(t *testing.T) {
	server := Smasher{}

	t.Run("it says hello world", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		server := Smasher{}
		server.ServeHTTP(response, request)
		if response.Body.String() != "Hello World" {
			t.Errorf("expected Hello World, got %q", response.Body)
		}
	})

	t.Run("it gets a response body from a url and puts it on the server response", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/?urls=bob", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		if response.Body.String() != "Bob" {
			t.Errorf("expected Bob, got %q", response.Body)
		}

	})

}
