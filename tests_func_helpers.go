package cruder

import (
	"database/sql"
	"net/http"
	"reflect"
)

func testSet() (out *cruderOptionHelper) {
	return &cruderOptionHelper{
		method: "",
		model:  reflect.TypeOf(struct{}{}),
		many:   reflect.SliceOf(reflect.TypeOf(struct{}{})),
		uri:    "",
	}
}

type testHttpResponseWriter struct {
	header     http.Header
	writeBytes []byte
	statusCode int
}

func (t *testHttpResponseWriter) Header() http.Header {
	return t.header
}

func (t *testHttpResponseWriter) Write(bytes []byte) (int, error) {
	t.writeBytes = append(t.writeBytes, bytes...)
	return len(bytes), nil
}

func (t *testHttpResponseWriter) WriteHeader(statusCode int) {
	t.statusCode = statusCode
}

func newTestHttpResponseWriter() (out *testHttpResponseWriter) {
	return &testHttpResponseWriter{
		header:     http.Header{},
		writeBytes: []byte{},
		statusCode: 200,
	}
}

type testDB struct{}

func (t *testDB) Offset(offset int) (tx IDB) {
	return t
}

func (t *testDB) Limit(limit int) (tx IDB) {
	return t
}

func (t *testDB) TableName(model interface{}) (out string) {
	return ""
}

func (t *testDB) Find(dest interface{}, conds ...interface{}) (tx IDB) {
	return t
}

func (t *testDB) Count(count *int64) (tx IDB) {
	return t
}

func (t *testDB) Error() (out error) {
	return nil
}

func (t *testDB) Unscoped() (tx IDB) {
	return t
}

func (t *testDB) Model(model interface{}) (tx IDB) {
	return t
}

func (t *testDB) Update(column string, value interface{}) (tx IDB) {
	return t
}

func (t *testDB) Debug() (tx IDB) {
	return t
}

func (t *testDB) Where(query interface{}, args ...interface{}) (tx IDB) {
	return t
}

func (t *testDB) Create(value interface{}) (tx IDB) {
	return t
}

func (t *testDB) Delete(value interface{}, conds ...interface{}) (tx IDB) {
	return t
}

func (t *testDB) Save(value interface{}) (tx IDB) {
	return t
}

func (t *testDB) DB() (*sql.DB, error) {
	return &sql.DB{}, nil
}

func тewTestDB() (out IDB) {
	return &testDB{}
}
