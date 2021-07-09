package cruder

import (
	"net/http"
)

// method меняет название методо для запроса
type method struct {
	method string
}

// Option имплементация интерфеса types.ICruderOption
func (o *method) Option(optionSetIn iCruderOptionHelper) {
	optionSetIn.SetMethod(o.method)
}

// Method меняет название методо для запроса
func Method(methodIn string) (out ICruderOption) {
	return &method{method: ValidMethod(methodIn)}
}

var (
	// Get завернутый в Method http.MethodGet
	Get = Method(http.MethodGet)
	// Head завернутый в Method http.MethodHead
	Head = Method(http.MethodHead)
	// Post завернутый в Method http.MethodPost
	Post = Method(http.MethodPost)
	// Put завернутый в Method http.MethodPut
	Put = Method(http.MethodPut)
	// Patch завернутый в Method http.MethodPatch
	Patch = Method(http.MethodPatch)
	// Delete завернутый в Method http.MethodDelete
	Delete = Method(http.MethodDelete)
	// Connect завернутый в Method http.MethodConnect
	Connect = Method(http.MethodConnect)
	// Options завернутый в Method http.MethodOptions
	Options = Method(http.MethodOptions)
	// Trace завернутый в Method http.MethodTrace
	Trace = Method(http.MethodTrace)
)
