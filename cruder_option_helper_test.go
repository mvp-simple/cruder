package cruder

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewOptionHelper(t *testing.T) {
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
		method    string
		model     interface{}
		oneModel  interface{}
		manyModel interface{}
		uri       string
	}{
		{
			name:      "1 test",
			method:    http.MethodGet,
			model:     test1{},
			oneModel:  test1{},
			manyModel: []test1{},
			uri:       "/1 test",
		},
		{
			name:      "2 test",
			method:    http.MethodGet,
			model:     test2{},
			oneModel:  test2{},
			manyModel: []test2{},
			uri:       "/2 test",
		},
		{
			name:      "test3",
			method:    http.MethodGet,
			model:     test3{},
			oneModel:  test3{},
			manyModel: []test3{},
			uri:       "/test3",
		},
		{
			name:      "test4",
			method:    http.MethodGet,
			model:     test4{},
			oneModel:  test4{},
			manyModel: []test4{},
			uri:       "/test4",
		},
		{
			name:      "test5",
			method:    http.MethodGet,
			model:     test5{},
			oneModel:  test5{},
			manyModel: []test5{},
			uri:       "/test5",
		},
		{
			name:      "test6",
			method:    http.MethodGet,
			model:     test6{},
			oneModel:  test6{},
			manyModel: []test6{},
			uri:       "/test6",
		},
		{
			name:      "test7",
			method:    http.MethodGet,
			model:     test7{},
			oneModel:  test7{},
			manyModel: []test7{},
			uri:       "/test7",
		},
		{
			name:      "test8",
			method:    http.MethodGet,
			model:     test8{},
			oneModel:  test8{},
			manyModel: []test8{},
			uri:       "/test8",
		},
		{
			name:      "test9",
			method:    http.MethodGet,
			model:     test9{},
			oneModel:  test9{},
			manyModel: []test9{},
			uri:       "/test9",
		},
		{
			name:      "test10",
			method:    http.MethodGet,
			model:     test10{},
			oneModel:  test10{},
			manyModel: []test10{},
			uri:       "/test10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			one := reflect.TypeOf(tt.model)
			slice := reflect.SliceOf(one)
			set := newOptionHelper(tt.method, one, slice, tt.uri)
			if !reflect.DeepEqual(reflect.TypeOf(tt.oneModel).String(), set.One().String()) {
				t.Errorf("One() = '%v', want '%v'", set.One().String(), reflect.TypeOf(tt.oneModel).String())
			}
			if !reflect.DeepEqual(reflect.TypeOf(tt.manyModel).String(), set.Many().String()) {
				t.Errorf("Many() = '%v', want '%v'", set.Many().String(), reflect.TypeOf(tt.manyModel).String())
			}
			if !reflect.DeepEqual(set.Method(), tt.method) {
				t.Errorf("Method() = '%v', want '%v'", reflect.TypeOf(set.Method()), reflect.TypeOf(tt.method))
			}
			if !reflect.DeepEqual(set.Uri(), tt.uri) {
				t.Errorf("Uri() = '%v', want '%v'", reflect.TypeOf(set.Uri()), reflect.TypeOf(tt.uri))
			}
		})
	}
}
