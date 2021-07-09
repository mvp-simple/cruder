package cruder

import (
	"encoding/json"
	"net/http"
)

var (
	statusInternalServerError = TRequestAnswer{
		Code:    http.StatusInternalServerError,
		Payload: "can not create json",
	}
	byteStatusInternalServerError, _ = json.Marshal(statusInternalServerError)
)

func CanNotCreateJson(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusInternalServerError)
	_, _ = writer.Write(byteStatusInternalServerError)
}

type writer struct {
	writer  http.ResponseWriter
	code    int
	payload interface{}
}

func (h *writer) SetCode(codeIn int) (out IWriter) {
	h.code = codeIn
	return h
}

func (h *writer) SetPayload(payloadIn interface{}) (out IWriter) {
	h.payload = payloadIn
	return h
}

func (h *writer) Write() (out IWriter) {
	if bytes, errMarshal := json.Marshal(TRequestAnswer{Code: h.code, Payload: h.payload}); errMarshal == nil {
		h.writer.WriteHeader(h.code)
		_, _ = h.writer.Write(bytes)
	} else {
		CanNotCreateJson(h.writer)
	}
	return h
}

func (h *writer) StatusForbidden() {
	h.SetCode(http.StatusForbidden).SetPayload("permission forbidden")
}

func (h *writer) StatusNotFound() {
	h.SetCode(http.StatusNotFound).SetPayload("url not found")
}

func (h *writer) StatusMethodNotAllowed() {
	h.SetCode(http.StatusMethodNotAllowed).SetPayload("method not allowed")
}

func NewWriter(w http.ResponseWriter) (out IWriter) {
	return &writer{writer: w, code: 200, payload: nil}
}
