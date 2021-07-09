package cruder

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type bodyManager struct {
	request *http.Request
	bytes   []byte
}

func (b *bodyManager) Bytes() (out []byte) {
	return b.bytes
}

func (b *bodyManager) Json(in interface{}) (errOut error) {
	return json.Unmarshal(b.Bytes(), in)
}

func NewBodyManager(r *http.Request) (out IBodyManager) {
	byteSlice, _ := ioutil.ReadAll(r.Body)
	return &bodyManager{
		request: r,
		bytes:   byteSlice,
	}
}
