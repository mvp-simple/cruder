package cruder

import (
	"net/http"
)

type reader struct {
	request *http.Request
	param   IParamManager
	body    IBodyManager
}

func NewReader(request *http.Request) IReader {
	return &reader{request: request, param: NewParamManager(request.URL), body: NewBodyManager(request)}
}

func (r *reader) Param() IParamManager {
	return r.param
}
func (r *reader) Body() IBodyManager {
	return r.body
}
