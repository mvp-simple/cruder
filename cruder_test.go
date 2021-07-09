package cruder

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type test1 struct{}
	type test2 struct{}
	type test3 struct{}
	type test4 struct{}
	type test5 struct{}
	type test6 struct{}
	type test7 struct{}
	type test8 struct{}
	type test9 struct{}
	type test10 struct{}
	tests := []struct {
		name  string
		model interface{}
	}{
		{
			name:  "test1",
			model: test1{},
		},
		{
			name:  "test2",
			model: test2{},
		},
		{
			name:  "test3",
			model: test3{},
		},
		{
			name:  "test4",
			model: test4{},
		},
		{
			name:  "test5",
			model: test5{},
		},
		{
			name:  "test6",
			model: test6{},
		},
		{
			name:  "test7",
			model: test7{},
		},
		{
			name:  "test8",
			model: test8{},
		},
		{
			name:  "test9",
			model: test9{},
		},
		{
			name:  "test10",
			model: test10{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.model, тewTestDB()); !reflect.DeepEqual(got.Model(), reflect.TypeOf(tt.model)) {
				t.Errorf("New().one = %v, want %v", got.Model(), reflect.TypeOf(tt.model))
			}
		})
	}
}
