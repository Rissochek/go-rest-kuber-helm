package server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	tt := []struct {
		name             string
		input            string
		expectedCode     int
		expectedResponse string
	}{
		{"test riss hello", "riss", http.StatusOK, fmt.Sprintf("Hello %v, u'r welcome!", "riss")},
		{"test empty name", "", http.StatusNotFound, "404 page not found\n"},
		{"test russian language", "Руслан", http.StatusOK, fmt.Sprintf("Hello %v, u'r welcome!", "Руслан")},
	}

	router := SetupRouter()

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest("GET", fmt.Sprintf("/hello/%v", test.input), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			response, err := io.ReadAll(w.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if w.Code != test.expectedCode || string(response) != test.expectedResponse {
				t.Errorf("expected code to be %v got %v", test.expectedCode, w.Code)
				t.Errorf("expected response to be %v got %v", test.expectedResponse, string(response))
			}
		})
	}
}
