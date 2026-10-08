package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKeyValidHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey abc123")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if key != "abc123" {
		t.Errorf("expected key abc123, got %q", key)
	}
}

func TestGetAPIKeyMissingHeader(t *testing.T) {
	headers := http.Header{}

	key, err := GetAPIKey(headers)

	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("expected missing-header error, got %v", err)
	}

	if key != "" {
		t.Errorf("expected empty key, got %q", key)
	}
}
