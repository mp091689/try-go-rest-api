package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi_hello(t *testing.T) {
	type fields struct {
		config *Config
	}
	tests := []struct {
		name     string
		route    string
		routeVar string
		fields   fields
	}{
		{"first test for hello", "/hello", "4", fields{&Config{Port: "3000"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := fmt.Sprintf("hello to you %s times", tt.routeVar)
			path := fmt.Sprintf("%s/%s", tt.route, tt.routeVar)

			req, err := http.NewRequest(http.MethodGet, path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()

			a := NewApi(tt.fields.config)
			a.ConfigureRouter()
			a.router.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Errorf("response code got: %d, want: %d", rec.Code, http.StatusOK)
			}

			gotBody := rec.Body.String()
			if gotBody != want {
				t.Errorf("response body got: %s, want: %s", gotBody, want)
			}
		})
	}
}
