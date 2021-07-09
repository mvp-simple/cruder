package cruder

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"
)

func TestNewReader(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		urlKey string
		urlStr string
		bytes  []byte
	}{
		{
			name:   "test 1",
			url:    "/kinder?one=1&one=2&one=3&one=4&one=0&b=1&b=0&b=true&b=false&bone=false",
			urlKey: "one",
			urlStr: "12340",
			bytes:  []byte("Hello world"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, errnewRequest := http.NewRequest(http.MethodPost, tt.url, bytes.NewReader(tt.bytes))
			if errnewRequest != nil {
				t.Fatal(errnewRequest)
			}
			testReader := NewReader(request)
			if !reflect.DeepEqual(testReader.Body().Bytes(), tt.bytes) {
				t.Errorf("NewReader().Body().Bytes() = %v, want %v", testReader.Body().Bytes(), tt.bytes)
			}
			str, errStr := testReader.Param().Str(tt.urlKey)
			if errStr != nil {
				t.Fatal(errStr)
			}
			if !reflect.DeepEqual(str, tt.urlStr) {
				t.Errorf("NewReader().Param().Str() = %v, want %v", str, tt.urlStr)
			}
		})
	}
}
