package cruder

import (
	"reflect"
	"testing"
)

func TestOneModelCreator(t *testing.T) {
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
		name      string
		creatorIn interface{}
		wantOut   interface{}
	}{
		{name: "1 test", creatorIn: test1{}, wantOut: test1{}},
		{name: "2 test", creatorIn: test2{}, wantOut: test2{}},
		{name: "3 test", creatorIn: test3{}, wantOut: test3{}},
		{name: "4 test", creatorIn: test4{}, wantOut: test4{}},
		{name: "5 test", creatorIn: test5{}, wantOut: test5{}},
		{name: "6 test", creatorIn: test6{}, wantOut: test6{}},
		{name: "7 test", creatorIn: test7{}, wantOut: test7{}},
		{name: "8 test", creatorIn: test8{}, wantOut: test8{}},
		{name: "9 test", creatorIn: test9{}, wantOut: test9{}},
		{name: "10 test", creatorIn: test10{}, wantOut: test10{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := testSet()
			OneModelCreator(tt.creatorIn).
				Option(set)
			if !reflect.DeepEqual(reflect.TypeOf(tt.wantOut).String(), set.model.String()) {
				t.Errorf("One() = '%v', want '%v'", set.model.String(), reflect.TypeOf(tt.wantOut).String())
			}
		})
	}
}
