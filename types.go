package cruder

import (
	"database/sql"
	"net/http"
	"reflect"
)

// structs

type ApiData struct {
	Func   func(ctx IContext)
	Method string
	Url    string
}

type TRequestAnswer struct {
	Code    int         `json:"code"`
	Payload interface{} `json:"payload"`
}

type ListResult struct {
	AllCount     int64       `json:"all_count"`
	List         interface{} `json:"list"`
	PageCount    int64       `json:"page_count"`
	RequestCount int64       `json:"request_count"`
}

// interfaces


type IBodyManager interface {
	Bytes() (out []byte)
	Json(in interface{}) (errOut error)
}

type IContext interface {
	HttpRequest() (out *http.Request)
	HttpResponseWriter() (out http.ResponseWriter)
	Reader() IReader
	SetUserID(in int64) (out IContext)
	SetUserUUID(in string) (out IContext)
	Writer() IWriter
	UserID() (out int64)
	UserUUID() (out string)
}

type IDB interface {
	Count(count *int64) (tx IDB)
	Create(value interface{}) (tx IDB)
	DB() (*sql.DB, error)
	Debug() (tx IDB)
	Delete(value interface{}, conds ...interface{}) (tx IDB)
	Error() (out error)
	Find(dest interface{}, conds ...interface{}) (tx IDB)
	Limit(limit int) (tx IDB)
	Model(model interface{}) (tx IDB)
	Offset(offset int) (tx IDB)
	Save(value interface{}) (tx IDB)
	TableName(model interface{}) (out string)
	Unscoped() (tx IDB)
	Update(column string, value interface{}) (tx IDB)
	Where(query interface{}, args ...interface{}) (tx IDB)
}

type IContextOption interface {
	Option(in IContext) (out IContextOption)
}

type ICruder interface {
	Api() (out []ApiData)
	Create(options ...ICruderOption) (out ICruder)
	Delete(options ...ICruderOption) (out ICruder)
	Get(options ...ICruderOption) (out ICruder)
	List(options ...ICruderOption) (out ICruder)
	Model() (out reflect.Type)
	Restore(options ...ICruderOption) (out ICruder)
	Update(options ...ICruderOption) (out ICruder)
}

type ICruderOption interface {
	Option(optionSetIn iCruderOptionHelper)
}

type iCruderOptionHelper interface {
	Many() (out reflect.Type)
	Method() (out string)
	One() (out reflect.Type)
	SetMethod(methodIn string)
	SetModel(model reflect.Type)
	SetModelSlice(modelSlice reflect.Type)
	SetUri(uriIn string)
	Uri() (out string)
}

type IParamManager interface {
	Bool(keyIn string) (out bool, errOut error)
	BoolSlice(keyIn string) (sliceOut []bool, errOut error)
	Float(keyIn string) (out float64, errOut error)
	FloatSlice(keyIn string) (sliceOut []float64, errOut error)
	Int(keyIn string) (out int64, errOut error)
	IntSlice(keyIn string) (sliceOut []int64, errOut error)
	Str(keyIn string) (out string, errOut error)
	StrSlice(keyIn string) (sliceOut []string, errOut error)
}

type IReader interface {
	Body() IBodyManager
	Param() IParamManager
}

type IRouter interface {
	Add(cruder ICruder) (out IRouter)
	AddApi(urlIn, methodIn string, funcIn func(ctx IContext)) (out IRouter)
	AddApiData(in ApiData) (out IRouter)
	ServeHTTP(writer http.ResponseWriter, request *http.Request)
	SetIdentity(in func(ctx IContext) (out bool)) (out IRouter)
}

type IRouterOption interface {
	Option(routerIn IRouter)
}

type IWriter interface {
	SetCode(codeIn int) (out IWriter)
	SetPayload(payloadIn interface{}) (out IWriter)
	StatusForbidden()
	StatusMethodNotAllowed()
	StatusNotFound()
	Write() (out IWriter)
}



