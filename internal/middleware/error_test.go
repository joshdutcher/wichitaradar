package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	tests := []struct {
		name           string
		handler        interface{}
		expectedStatus int
		expectedBody   string
		prodBody      string
	}{
		{
			name: "standard http.Handler",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("OK"))
			}),
			expectedStatus: http.StatusOK,
			expectedBody:   "OK",
			prodBody:      "OK",
		},
		{
			name: "HandlerFunc returning error",
			handler: HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
				return InternalError(fmt.Errorf("test error"))
			}),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Internal Server Error\n\nError: Internal error: test error\n",
			prodBody:      "Internal Server Error\n",
		},
		{
			name: "func returning error",
			handler: func(w http.ResponseWriter, r *http.Request) error {
				return BadRequestError(fmt.Errorf("invalid input"), "Invalid input")
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid input\n\nError: Bad request: Invalid input: invalid input\n",
			prodBody:      "Invalid input\n",
		},
		{
			name: "unsupported handler type",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "",
			prodBody:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test production mode
			req := httptest.NewRequest("GET", "/", nil)
			req.Host = "example.com" // Production mode
			rr := httptest.NewRecorder()

			// Set ENV for production test
			os.Setenv("ENV", "production")
			defer os.Unsetenv("ENV")

			handler := ErrorHandler(tt.handler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if body := rr.Body.String(); body != tt.prodBody {
				t.Errorf("handler returned unexpected body in production: got %q want %q",
					body, tt.prodBody)
			}

			// Test development mode
			req = httptest.NewRequest("GET", "/", nil)
			req.Host = "localhost" // Development mode
			rr = httptest.NewRecorder()

			// Ensure ENV is not set for development test
			os.Unsetenv("ENV")

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code in dev mode: got %v want %v",
					status, tt.expectedStatus)
			}

			if body := rr.Body.String(); body != tt.expectedBody {
				t.Errorf("handler returned unexpected body in dev mode: got %q want %q",
					body, tt.expectedBody)
			}
		})
	}
}