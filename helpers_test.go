package cruder

import (
	"net/http"
	"testing"
)

func TestValidMethod(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		wantMethod string
	}{
		{name: http.MethodGet, method: http.MethodGet, wantMethod: http.MethodGet},
		{name: http.MethodHead, method: http.MethodHead, wantMethod: http.MethodHead},
		{name: http.MethodPost, method: http.MethodPost, wantMethod: http.MethodPost},
		{name: http.MethodPut, method: http.MethodPut, wantMethod: http.MethodPut},
		{name: http.MethodPatch, method: http.MethodPatch, wantMethod: http.MethodPatch},
		{name: http.MethodDelete, method: http.MethodDelete, wantMethod: http.MethodDelete},
		{name: http.MethodConnect, method: http.MethodConnect, wantMethod: http.MethodConnect},
		{name: http.MethodOptions, method: http.MethodOptions, wantMethod: http.MethodOptions},
		{name: http.MethodTrace, method: http.MethodTrace, wantMethod: http.MethodTrace},
		{name: "any", method: "any", wantMethod: http.MethodGet},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := ValidMethod(tt.method); gotOut != tt.wantMethod {
				t.Errorf("Method() = %v, want %v", gotOut, tt.wantMethod)
			}
		})
	}
}
