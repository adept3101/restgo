package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rest/api"
	"testing"
)

// var db = api.ConnDB()
//
// var app = &api.App{
// 	DB: db,
// }

func TestGetUser(t *testing.T) {
	// expect := 1
	req := httptest.NewRequest(http.MethodGet, "/user/{1}", nil)
	rec := httptest.NewRecorder()

	app.GetUsersHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf(
			"Expected status %d, got %d",
			http.StatusOK,
			rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}

	var response api.UsersResponse

	err := json.NewDecoder(rec.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Status != 200 {
		t.Errorf("Expected status 200, got %d", response.Status)
	}

	if len(response.Data) == 0 {
		t.Error("Expected user, got 0")
	}
}
