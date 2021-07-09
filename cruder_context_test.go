package cruder

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"
)

func TestNewContext(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		urlKey   string
		urlValue string
		body     []byte
	}{
		{
			name:     "test 1",
			url:      "?param1=terry",
			urlKey:   "param1",
			urlValue: "terry",
			body:     []byte("Hello world"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseWriter := newTestHttpResponseWriter()
			request, errNewRequest := http.NewRequest(http.MethodPost, tt.url, bytes.NewReader(tt.body))
			if errNewRequest != nil {
				t.Fatal(errNewRequest)
			}
			ctx := NewContext(responseWriter, request)
			if !reflect.DeepEqual(ctx.HttpResponseWriter(), responseWriter) {
				t.Errorf("ctx.HttpResponseWriter() = %v, want %v", ctx.HttpResponseWriter(), responseWriter)
			}
			if !reflect.DeepEqual(ctx.HttpRequest(), request) {
				t.Errorf("ctx.HttpResponseWriter() = %v, want %v", ctx.HttpResponseWriter(), responseWriter)
			}
			if !reflect.DeepEqual(ctx.UserID(), int64(0)) {
				t.Errorf("ctx.UserID() = %v, want %v", ctx.UserID(), 0)
			}
			if !reflect.DeepEqual(ctx.UserUUID(), "") {
				t.Errorf("ctx.UserUUID() = %v, want %v", ctx.UserUUID(), "")
			}
			if !reflect.DeepEqual(ctx.Reader().Body().Bytes(), tt.body) {
				t.Errorf("ctx.Reader().Body().Bytes() = %v, want %v", ctx.Reader().Body().Bytes(), tt.body)
			}
			urlValue, errUrlValue := ctx.Reader().Param().Str(tt.urlKey)
			if errUrlValue != nil {
				t.Fatal(errUrlValue)
			}
			if !reflect.DeepEqual(urlValue, tt.urlValue) {
				t.Errorf("ctx.Reader().Param().Str(%s) = %v, want %v",tt.urlKey, urlValue, tt.urlValue)
			}
		})
	}
}

func TestContextUUID(t *testing.T) {
	tests := []struct {
		name string
		data string
	}{
		{
			name: "test 1",
			data: "1",
		},
		{
			name: "test 2",
			data: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseWriter := newTestHttpResponseWriter()
			request, errNewRequest := http.NewRequest(http.MethodPost, "", bytes.NewReader(nil))
			if errNewRequest != nil {
				t.Fatal(errNewRequest)
			}
			ctx := NewContext(responseWriter, request, ContextUUID(tt.data))
			if !reflect.DeepEqual(ctx.UserUUID(), tt.data) {
				t.Errorf("ContextUUID() = %v, want %v", ctx.UserUUID(), tt.data)
			}
		})
	}
}

func TestContextID(t *testing.T) {
	tests := []struct {
		name string
		data int64
	}{
		{
			name: "test 1",
			data: 1,
		},
		{
			name: "test 2",
			data: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseWriter := newTestHttpResponseWriter()
			request, errnewRequest := http.NewRequest(http.MethodPost, "", bytes.NewReader(nil))
			if errnewRequest != nil {
				t.Fatal(errnewRequest)
			}
			ctx := NewContext(responseWriter, request, ContextID(tt.data))
			if !reflect.DeepEqual(ctx.UserID(), tt.data) {
				t.Errorf("ContextUUID() = %v, want %v", ctx.UserID(), tt.data)
			}
		})
	}
}
