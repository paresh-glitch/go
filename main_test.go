package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	http.HandlerFunc(homeHandler).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", rr.Code)
	}

	var resp Response
	json.NewDecoder(rr.Body).Decode(&resp)

	if resp.Message == "" {
		t.Error("message should not be empty")
	}
	if resp.Language != "Go" {
		t.Errorf("expected Go got %s", resp.Language)
	}
}

func TestHealthHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()
	http.HandlerFunc(healthHandler).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", rr.Code)
	}

	var resp HealthResponse
	json.NewDecoder(rr.Body).Decode(&resp)

	if resp.Status != "healthy" {
		t.Errorf("expected healthy got %s", resp.Status)
	}
}

func TestUsersHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/users", nil)
	rr := httptest.NewRecorder()
	http.HandlerFunc(usersHandler).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 got %d", rr.Code)
	}

	var users []User
	json.NewDecoder(rr.Body).Decode(&users)

	if len(users) != 3 {
		t.Errorf("expected 3 users got %d", len(users))
	}
}
