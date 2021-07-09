package cruder

import (
	"net/http"
)

// ValidMethod проверяет и выдает скорректированный метод
func ValidMethod(methodIn string) (out string) {
	switch methodIn {
	case http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace:
	default:
		methodIn = http.MethodGet
	}
	return methodIn
}
