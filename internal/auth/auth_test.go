package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("Successful API key extraction", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey test-api-key")

		apiKey, err := GetAPIKey(headers)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if apiKey != "test-api-key" {
			t.Errorf("Expected API key 'test-api-key', got '%s'", apiKey)
		}
	})

	t.Run("No authorization header", func(t *testing.T) {
		headers := http.Header{}

		_, err := GetAPIKey(headers)
		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("Expected ErrNoAuthHeaderIncluded, got %v", err)
		}
	})
}
