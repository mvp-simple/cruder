package cruder

import (
	"reflect"
	"testing"
)

func TestManyModelCreator(t *testing.T) {
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
		name    string
		model   interface{}
		wantOut interface{}
	}{
		{name: "1 test", model: []test1{}, wantOut: []test1{}},
		{name: "2 test", model: []test2{}, wantOut: []test2{}},
		{name: "3 test", model: []test3{}, wantOut: []test3{}},
		{name: "4 test", model: []test4{}, wantOut: []test4{}},
		{name: "5 test", model: []test5{}, wantOut: []test5{}},
		{name: "6 test", model: []test6{}, wantOut: []test6{}},
		{name: "7 test", model: []test7{}, wantOut: []test7{}},
		{name: "8 test", model: []test8{}, wantOut: []test8{}},
		{name: "9 test", model: []test9{}, wantOut: []test9{}},
		{name: "10 test", model: []test10{}, wantOut: []test10{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := testSet()
			ManyModelCreator(tt.model).
				Option(set)
			if !reflect.DeepEqual(reflect.TypeOf(tt.wantOut).String(), set.many.String()) {
				t.Errorf("One() = '%v', want '%v'", set.many.String(), reflect.TypeOf(tt.wantOut).String())
			}
		})
	}
}
