package server

import (
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	tt := []struct {
		name                string
		expectedContentType string
	}{
		{"test content type", "text/plain; charset=utf-8"},
	}

	router := SetupRouter()

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			r := httptest.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)
			response := w.Result()
			contentType := response.Header.Get("Content-Type")
			if contentType != test.expectedContentType {
				t.Errorf("expected content type to be %v got %v", test.expectedContentType, contentType)
			}
		})
	}
}
