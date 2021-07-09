package cruder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCanNotCreateJson(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				CanNotCreateJson(w)
			}))
			defer ts.Close()
			res, errGet := http.Get(ts.URL)
			if errGet != nil {
				t.Fatal(errGet)
			}
			defer res.Body.Close()
			if res.StatusCode != http.StatusInternalServerError {
				t.Errorf("StatusCode = %v, want %v", res.StatusCode, http.StatusInternalServerError)
			}
			bytes, errReadAll := io.ReadAll(res.Body)
			if errReadAll != nil {
				t.Fatal(errReadAll)
			}
			var answer TRequestAnswer
			errUnmarshal := json.Unmarshal(bytes, &answer)
			if errUnmarshal != nil {
				t.Fatal(errUnmarshal)
			}
			if answer.Code != statusInternalServerError.Code {
				t.Errorf("answer.Code = %v, want %v", answer.Code, statusInternalServerError.Code)
			}
			if fmt.Sprint(answer.Payload) != fmt.Sprint(statusInternalServerError.Payload) {
				t.Errorf("answer.Code = %s, want %s", fmt.Sprint(answer.Payload), fmt.Sprint(statusInternalServerError.Payload))
			}
		}()
	})
}

func TestNewWriter(t *testing.T) {
	tests := []struct {
		name    string
		payload interface{}
		code    int
	}{
		{
			name:    "test1",
			payload: "test 1",
			code:    201,
		},
		{
			name:    "test2",
			payload: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseWriter := newTestHttpResponseWriter()
			writer := NewWriter(responseWriter)
			writer.SetCode(tt.code).SetPayload(tt.payload).Write()
			if !reflect.DeepEqual(responseWriter.statusCode, tt.code) {
				t.Errorf("statusCode = %v, want %v", responseWriter.statusCode, tt.code)
			}
			var answer TRequestAnswer
			errUnmarshal := json.Unmarshal(responseWriter.writeBytes, &answer)
			if errUnmarshal != nil {
				t.Fatal(errUnmarshal)
			}
			if !reflect.DeepEqual(answer.Code, tt.code) {
				t.Errorf("Code = %v, want %v", answer.Code, tt.code)
			}
			if !reflect.DeepEqual(fmt.Sprint(answer.Payload), fmt.Sprint(tt.payload)) {
				t.Errorf("Payload = %v, want %v", answer.Payload, tt.payload)
			}
		})
	}
}
