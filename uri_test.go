package cruder

import (
	"reflect"
	"testing"
)

func TestUri(t *testing.T) {
	tests := []struct {
		name  string
		uriIn string
	}{
		{name: "1 test", uriIn: "/uri/ppp1"},
		{name: "2 test", uriIn: "/uri/ppp2"},
		{name: "3 test", uriIn: "/uri/ppp3"},
		{name: "4 test", uriIn: "/uri/ppp4"},
		{name: "5 test", uriIn: "/uri/ppp5"},
		{name: "6 test", uriIn: "/uri/ppp6"},
		{name: "7 test", uriIn: "/uri/ppp7"},
		{name: "8 test", uriIn: "/uri/ppp8"},
		{name: "9 test", uriIn: "/uri/ppp9"},
		{name: "10 test", uriIn: "/uri/ppp10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := testSet()
			Uri(tt.uriIn).
				Option(set)
			if !reflect.DeepEqual(tt.uriIn, set.uri) {
				t.Errorf("Uri() = %v, want %v", tt.uriIn, set.uri)
			}
		})
	}
}
