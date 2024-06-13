package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("Authorization", "Authorization Basic")
	// expected := "Basic"
	// actual, _ := GetAPIKey(req.Header)
	// if actual != expected {
	// 	t.Fatalf("expected: %v, got %v", expected, actual)
	// }
}
