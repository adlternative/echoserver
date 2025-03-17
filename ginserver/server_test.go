package ginserver

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
}

func TestEchoHandlerGin(t *testing.T) {
	// Create a router with the echo handler
	router := SetupRouter()

	// Test cases
	testCases := []struct {
		name           string
		method         string
		path           string
		body           string
		contentType    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "POST with JSON",
			method:         "POST",
			path:           "/test",
			body:           `{"message": "Hello, Echo!"}`,
			contentType:    "application/json",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message": "Hello, Echo!"}`,
		},
		{
			name:           "GET with text",
			method:         "GET",
			path:           "/",
			body:           "Hello, Echo!",
			contentType:    "text/plain",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, Echo!",
		},
		{
			name:           "PUT with empty body",
			method:         "PUT",
			path:           "/empty",
			body:           "",
			contentType:    "text/plain",
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:           "DELETE with XML",
			method:         "DELETE",
			path:           "/xml",
			body:           "<data>test</data>",
			contentType:    "application/xml",
			expectedStatus: http.StatusOK,
			expectedBody:   "<data>test</data>",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request with the test case data
			req, err := http.NewRequest(tc.method, tc.path, bytes.NewBufferString(tc.body))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Set Content-Type header if provided
			if tc.contentType != "" {
				req.Header.Set("Content-Type", tc.contentType)
			}

			// Create a response recorder
			w := httptest.NewRecorder()

			// Serve the request
			router.ServeHTTP(w, req)

			// Check status code
			if w.Code != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, w.Code)
			}

			// Check response body
			if strings.TrimSpace(w.Body.String()) != strings.TrimSpace(tc.expectedBody) {
				t.Errorf("Expected body %q, got %q", tc.expectedBody, w.Body.String())
			}

			// Check Content-Type header
			if tc.contentType != "" && w.Header().Get("Content-Type") != tc.contentType {
				t.Errorf("Expected Content-Type %q, got %q", tc.contentType, w.Header().Get("Content-Type"))
			}
		})
	}
}

func TestSetupRouter(t *testing.T) {
	router := SetupRouter()
	if router == nil {
		t.Error("Expected router to be initialized, got nil")
	}
}
