package cruder

import (
	"net/url"
	"reflect"
	"testing"
)

func TestNewParamManager(t *testing.T) {
	tests := []struct {
		name          string
		uri           string
		StrSliceKey   string
		StrSlice      []string
		StrKey        string
		Str           string
		IntSliceKey   string
		IntSlice      []int64
		IntKey        string
		Int           int64
		BoolSliceKey  string
		BoolSlice     []bool
		BoolKey       string
		Bool          bool
		FloatSliceKey string
		FloatSlice    []float64
		FloatKey      string
		Float         float64
	}{
		{
			name:          "test 1",
			uri:           "/kinder?one=1&one=2&one=3&one=4&one=0&b=1&b=0&b=true&b=false&bone=false",
			StrSliceKey:   "one",
			StrSlice:      []string{"1", "2", "3", "4", "0"},
			StrKey:        "one",
			Str:           "12340",
			IntSliceKey:   "one",
			IntSlice:      []int64{1, 2, 3, 4, 0},
			IntKey:        "one",
			Int:           12340,
			BoolSliceKey:  "b",
			BoolSlice:     []bool{true, false, true, false},
			BoolKey:       "bone",
			Bool:          false,
			FloatSliceKey: "one",
			FloatSlice:    []float64{1, 2, 3, 4, 0},
			FloatKey:      "one",
			Float:         12340,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uriParsed, errParse := url.Parse(tt.uri)
			if errParse != nil {
				t.Fatal(errParse)
			}
			manager := NewParamManager(uriParsed)
			{
				data, err := manager.Str(tt.StrKey)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(data, tt.Str) {
					t.Errorf("Str() = %s, want %s", data, tt.Str)
				}
			}
			{
				data, err := manager.StrSlice(tt.StrSliceKey)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(data, tt.StrSlice) {
					t.Errorf("StrSlice() = %s, want %s", data, tt.StrSlice)
				}
			}
			{
				data, err := manager.Int(tt.IntKey)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(data, tt.Int) {
					t.Errorf("Int() = %v, want %v", data, tt.Int)
				}
			}
			{
				data, err := manager.IntSlice(tt.IntSliceKey)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(data, tt.IntSlice) {
					t.Errorf("IntSlice() = %v, want %v", data, tt.IntSlice)
				}
			}
			{
				data, err := manager.Float(tt.FloatKey)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(data, tt.Float) {
					t.Errorf("Float() = %v, want %v", data, tt.Float)
				}
			}
			{
				data, err := manager.FloatSlice(tt.FloatSliceKey)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(data, tt.FloatSlice) {
					t.Errorf("FloatSlice() = %v, want %v", data, tt.FloatSlice)
				}
			}
			{
				data, err := manager.Bool(tt.BoolKey)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(data, tt.Bool) {
					t.Errorf("Bool() = %v, want %v", data, tt.Bool)
				}
			}
			{
				data, err := manager.BoolSlice(tt.BoolSliceKey)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(data, tt.BoolSlice) {
					t.Errorf("BoolSlice() = %v, want %v", data, tt.BoolSlice)
				}
			}
		})
	}
}
