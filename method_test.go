package cruder

import (
	"net/http"
	"reflect"
	"testing"
)

func TestMethod(t *testing.T) {
	type testData struct {
		methodIn    string
		methodCheck string
		wantOut     cruderOptionHelper
	}

	tests := []testData{
		testData{methodIn: http.MethodGet, methodCheck: http.MethodGet, wantOut: cruderOptionHelper{method: http.MethodGet}},
		testData{methodIn: http.MethodHead, methodCheck: http.MethodHead, wantOut: cruderOptionHelper{method: http.MethodHead}},
		testData{methodIn: http.MethodPost, methodCheck: http.MethodPost, wantOut: cruderOptionHelper{method: http.MethodPost}},
		testData{methodIn: http.MethodPut, methodCheck: http.MethodPut, wantOut: cruderOptionHelper{method: http.MethodPut}},
		testData{methodIn: http.MethodPatch, methodCheck: http.MethodPatch, wantOut: cruderOptionHelper{method: http.MethodPatch}},
		testData{methodIn: http.MethodDelete, methodCheck: http.MethodDelete, wantOut: cruderOptionHelper{method: http.MethodDelete}},
		testData{methodIn: http.MethodConnect, methodCheck: http.MethodConnect, wantOut: cruderOptionHelper{method: http.MethodConnect}},
		testData{methodIn: http.MethodOptions, methodCheck: http.MethodOptions, wantOut: cruderOptionHelper{method: http.MethodOptions}},
		testData{methodIn: http.MethodTrace, methodCheck: http.MethodTrace, wantOut: cruderOptionHelper{method: http.MethodTrace}},
		testData{methodIn: "any", methodCheck: http.MethodGet, wantOut: cruderOptionHelper{method: http.MethodGet}},
	}
	for _, tt := range tests {
		t.Run(tt.methodIn, func(t *testing.T) {
			setOption := testSet()
			option := Method(tt.methodIn)
			option.Option(setOption)
			if !reflect.DeepEqual(tt.methodCheck, tt.wantOut.method) {
				t.Errorf("Method() = %v, want %v", tt.methodCheck, tt.wantOut)
			}
		})
	}
}

func TestVars(t *testing.T) {
	type testData struct {
		name        string
		method      ICruderOption
		methodCheck string
	}
	tests := []testData{
		testData{name: http.MethodGet, method: Get, methodCheck: http.MethodGet},
		testData{name: http.MethodHead, method: Head, methodCheck: http.MethodHead},
		testData{name: http.MethodPost, method: Post, methodCheck: http.MethodPost},
		testData{name: http.MethodPut, method: Put, methodCheck: http.MethodPut},
		testData{name: http.MethodPatch, method: Patch, methodCheck: http.MethodPatch},
		testData{name: http.MethodDelete, method: Delete, methodCheck: http.MethodDelete},
		testData{name: http.MethodConnect, method: Connect, methodCheck: http.MethodConnect},
		testData{name: http.MethodOptions, method: Options, methodCheck: http.MethodOptions},
		testData{name: http.MethodTrace, method: Trace, methodCheck: http.MethodTrace},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := testSet()
			tt.method.Option(set)
			if !reflect.DeepEqual(tt.methodCheck, set.method) {
				t.Errorf("%s != %v, want %s", tt.name, tt.methodCheck, tt.name)
			}
		})
	}
}
