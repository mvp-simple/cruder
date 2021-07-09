package cruder

import (
	"net/http"
)

type router struct {
	handlers map[string]map[string]func(ctx IContext)
	identity func(ctx IContext) (out bool)
}

func (r *router) SetIdentity(in func(ctx IContext) (out bool)) (out IRouter) {
	r.identity = in
	return r
}

func (r *router) AddApi(urlIn, methodIn string, funcIn func(ctx IContext)) (out IRouter) {
	if _, ok := r.handlers[urlIn]; !ok {
		r.handlers[urlIn] = make(map[string]func(ctx IContext))
	}
	r.handlers[urlIn][methodIn] = funcIn
	return r
}

func (r *router) AddApiData(in ApiData) (out IRouter) {
	return r.AddApi(in.Url, in.Method, in.Func)
}

func (r *router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := NewContext(writer, request)
	defer func() {
		if panicInterface := recover(); panicInterface != nil {
			ctx.Writer().SetCode(http.StatusInternalServerError).SetPayload(panicInterface)
		} else {
			ctx.Writer().Write()
		}
	}()

	if !r.identity(ctx) {
		ctx.Writer().StatusForbidden()
		return
	}

	data, okURL := r.handlers[request.URL.Path]
	if !okURL {
		ctx.Writer().StatusNotFound()
		return
	}

	handler, okMethod := data[request.Method]
	if !okMethod {
		ctx.Writer().StatusMethodNotAllowed()
		return
	}

	handler(ctx)
}

func (r *router) Add(cruder ICruder) (out IRouter) {
	for _, data := range cruder.Api() {
		r.AddApiData(data)
	}
	return r
}

func NewRouter(options ...IRouterOption) (out IRouter) {
	r := router{
		handlers: make(map[string]map[string]func(ctx IContext)),
		identity: func(ctx IContext) (out bool) { return true },
	}
	for _, option := range options {
		option.Option(&r)
	}
	return &r
}
