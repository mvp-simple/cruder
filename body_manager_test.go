package cruder

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestNewBodyManager(t *testing.T) {
	tests := []struct {
		name  string
		bytes []byte
		data  TRequestAnswer
	}{
		{
			name:  "test 1",
			bytes: byteStatusInternalServerError,
			data:  statusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, errNewRequest := http.NewRequest(http.MethodPost, "", bytes.NewReader(tt.bytes))
			if errNewRequest != nil {
				t.Fatal(errNewRequest)
			}
			manager := NewBodyManager(request)
			if !reflect.DeepEqual(manager.Bytes(), tt.bytes) {
				t.Errorf("NewBodyManager().Bytes() = %v, want %v", manager.Bytes(), tt.bytes)
			}
			var model TRequestAnswer
			errUnmarshal := manager.Json(&model)
			if errUnmarshal != nil {
				t.Fatal(errUnmarshal)
			}
			if !reflect.DeepEqual(model, tt.data) {
				t.Errorf("NewBodyManager().Json() = %v, want %v", model, tt.data)
			}
		})
	}
}

func Test_bodyManager_Bytes(t *testing.T) {
	tests := []struct {
		name  string
		bytes []byte
	}{
		{
			name:  "test 1",
			bytes: []byte("test 1"),
		},
		{
			name:  "test 2",
			bytes: []byte("test 2"),
		},
		{
			name:  "test 3",
			bytes: []byte("test 3"),
		},
		{
			name:  "test 4",
			bytes: []byte("test 4"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bodyManager{
				request: &http.Request{},
				bytes:   tt.bytes,
			}
			if gotOut := b.Bytes(); !reflect.DeepEqual(gotOut, tt.bytes) {
				t.Errorf("Bytes() = %v, want %v", gotOut, tt.bytes)
			}
		})
	}
}

func Test_bodyManager_Json(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		out  interface{}
	}{
		{
			name: "test0",
			in:   "string",
			out:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			byteSlice, errMarshal := json.Marshal(&tt.in)
			if errMarshal != nil {
				t.Fatal(errMarshal)
			}
			request, errNewRequest := http.NewRequest(http.MethodGet, "", bytes.NewReader(byteSlice))
			if errNewRequest != nil {
				t.Fatal(errNewRequest)
			}
			b := &bodyManager{
				request: request,
				bytes:   byteSlice,
			}

			if _ = b.Json(&tt.out); !reflect.DeepEqual(tt.out, tt.in) {
				t.Errorf("Json() out interface %v, want  %v", tt.out, tt.in)
			}
		})
	}
}

