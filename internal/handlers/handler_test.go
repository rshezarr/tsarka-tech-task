package handlers

import (
	"net/http"
	"net/http/httptest"
	"rest-api/internal/service"
	"strings"
	"testing"
)

func TestFindHandler(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "aaabbbccc112233",
			expected: "abc123",
		},
		{
			input:    "banana\nbandana",
			expected: "ban\nd",
		},
		{
			input:    "   \n",
			expected: " \n",
		},
		{
			input:    "313aca",
			expected: "31ac",
		},
		{
			input:    "AAaaBBbbCCcc",
			expected: "AaBbCc",
		},
	}

	for _, testCase := range testCases {
		// Create a new request
		req, err := http.NewRequest("POST", "/rest/substr/find", strings.NewReader(testCase.input))
		if err != nil {
			t.Fatalf("error while sending request: %e", err)
		}

		// Set content type
		req.Header.Set("Content-Type", "text/plain")

		// Create a new response recorder
		res := httptest.NewRecorder()

		// Create a new handler and serve the request
		h := &Handler{
			mux:     http.NewServeMux(),
			service: service.NewService(),
		}

		h.findHandler(res, req)

		// Check the response status code
		if res.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)
		}

		if res.Body.String() != testCase.expected {
			t.Errorf("Expected response body '%s', but got '%s'", testCase.expected, res.Body.String())
		}
	}
}

func TestEmailCheckHandler(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "test@example.com\ntest2@example.com",
			expected: `{"emails":["test@example.com","test2@example.com"]}`,
		},
		{
			input:    "test@example.com   @example.com",
			expected: `{"emails":["test@example.com"]}`,
		},
		{
			input:    "test@example.com   test2@example.com",
			expected: `{"emails":["test@example.com","test2@example.com"]}`,
		},
		{
			input:    "test@example.com test2@example.com\ntest3@example.com",
			expected: `{"emails":["test@example.com","test2@example.com","test3@example.com"]}`,
		},
	}

	for _, testCase := range testCases {
		// Create a new request
		req, err := http.NewRequest("POST", "/check-emails", strings.NewReader(testCase.input))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "text/plain")

		// Create a new response recorder
		res := httptest.NewRecorder()

		// Create a new handler and serve the request
		h := &Handler{
			mux:     http.NewServeMux(),
			service: service.NewService(),
		}

		h.checkEmailHandler(res, req)

		// Check the response status code
		if res.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)
		}

		// Check the response body
		if res.Body.String() != testCase.expected {
			t.Errorf("Expected response body '%s', but got '%s'", testCase.expected, res.Body.String())
		}
	}
}
