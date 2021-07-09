package cruder

import "net/http"

type cruderContext struct {
	w      http.ResponseWriter
	r      *http.Request
	writer IWriter
	reader IReader
	id     int64
	uuid   string
}

func (c *cruderContext) HttpResponseWriter() (out http.ResponseWriter) {
	return c.w
}

func (c *cruderContext) HttpRequest() (out *http.Request) {
	return c.r
}

func (c *cruderContext) UserID() (out int64) {
	return c.id
}

func (c *cruderContext) SetUserID(in int64) (out IContext) {
	c.id = in
	return c
}

func (c *cruderContext) UserUUID() (out string) {
	return c.uuid
}

func (c *cruderContext) SetUserUUID(in string) (out IContext) {
	c.uuid = in
	return c
}

func (c *cruderContext) Writer() IWriter {
	return c.writer
}

func (c *cruderContext) Reader() IReader {
	return c.reader
}

func NewContext(w http.ResponseWriter, r *http.Request, options ...IContextOption) (out IContext) {
	out = &cruderContext{
		w:      w,
		r:      r,
		id:     0,
		uuid:   "",
		reader: NewReader(r),
		writer: NewWriter(w),
	}
	for _, option := range options {
		option.Option(out)
	}
	return
}

type contextOptionID struct {
	data int64
}

func (o *contextOptionID) Option(in IContext) (out IContextOption) {
	in.SetUserID(o.data)
	return o
}

func ContextID(data int64) (out IContextOption) {
	return &contextOptionID{data: data}
}

type contextOptionUUID struct {
	data string
}

func (o *contextOptionUUID) Option(in IContext) (out IContextOption) {
	in.SetUserUUID(o.data)
	return o
}

func ContextUUID(data string) (out IContextOption) {
	return &contextOptionUUID{data: data}
}
