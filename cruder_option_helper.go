package cruder

import "reflect"

type cruderOptionHelper struct {
	method string
	model  reflect.Type
	many   reflect.Type
	uri    string
}

func newOptionHelper(methodIn string, model reflect.Type, modelSlice reflect.Type, uriIn string) (out iCruderOptionHelper) {
	return &cruderOptionHelper{method: ValidMethod(methodIn), model: model, many: modelSlice, uri: uriIn}
}

func (s *cruderOptionHelper) SetMethod(methodIn string) {
	s.method = methodIn
}

func (s *cruderOptionHelper) Method() (out string) {
	return s.method
}

func (s *cruderOptionHelper) SetModel(modelIn reflect.Type) {
	s.model = modelIn
}

func (s *cruderOptionHelper) One() (out reflect.Type) {
	return s.model
}

func (s *cruderOptionHelper) SetModelSlice(modelSliceIn reflect.Type) {
	s.many = modelSliceIn
}

func (s *cruderOptionHelper) Many() (out reflect.Type) {
	return s.many
}

func (s *cruderOptionHelper) SetUri(uriIn string) {
	s.uri = uriIn
}

func (s *cruderOptionHelper) Uri() (out string) {
	return s.uri
}
