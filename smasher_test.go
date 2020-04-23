package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSmasher(t *testing.T) {
	t.Run("it says hello world", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		server := Smasher{}
		server.ServeHTTP(response, request)
		if response.Body.String() != "Hello World" {
			t.Errorf("expected Hello World, got %q", response.Body)
		}
	})
}
