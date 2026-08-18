package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rest/api"
	"testing"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	record := httptest.NewRecorder()

	api.HealthHandler(record, req)

	if record.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, record.Code)
	}

	contentType := record.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var response api.Response

	err := json.NewDecoder(record.Body).Decode(&response)

	if err != nil {
		t.Fatalf("Error to decode response: %v", err)
	}

	if response.Message != "API is helthy" {
		t.Errorf("expected msg %q, got %q", "API is helthy", response.Message)
	}

	if response.Status != 200 {
		t.Errorf("expected status 200, got %d", response.Status)
	}
}
