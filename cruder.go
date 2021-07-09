package cruder

import (
	"net/http"
	"reflect"
	"strings"
)

type cruder struct {
	api        []ApiData
	one        reflect.Type
	many       reflect.Type
	tableName  string
	db         IDB
	apiList    string
	apiGet     string
	apiCreate  string
	apiDelete  string
	apiRestore string
	apiChange  string
}

func (c *cruder) modelUrlMethod(methodDefault string, modelDefault reflect.Type, modelSliceDefault reflect.Type,
	uriDefault string, options ...ICruderOption) (model reflect.Type, modelSlice reflect.Type, url, method string) {
	set := newOptionHelper(methodDefault, modelDefault, modelSliceDefault, uriDefault)
	for _, option := range options {
		option.Option(set)
	}
	model, modelSlice, url, method = set.One(), set.Many(), set.Uri(), set.Method()
	return
}

func (c *cruder) Api() (out []ApiData) {
	return c.api
}

func (c *cruder) Model() (out reflect.Type) {
	return c.one
}

func (c *cruder) List(options ...ICruderOption) (out ICruder) {
	_, many, url, methodUrl := c.modelUrlMethod(http.MethodGet, c.one, c.many, c.apiList, options...)
	c.api = append(c.api,
		ApiData{
			Url:    url,
			Method: methodUrl,
			Func: func(ctx IContext) {
				count, errCount := ctx.Reader().Param().Int("count")
				if errCount != nil || count == 0 {
					count = 100
				}
				page, errPage := ctx.Reader().Param().Int("page")
				if errPage != nil {
					page = 0
				}

				var (
					data      = reflect.New(many).Interface()
					listCount int64
				)
				if err := c.db.Offset(int(count * page)).Limit(int(count)).Find(data).Count(&listCount).Error(); err == nil {
					PageCount := listCount / count
					if listCount%count != 0 {
						PageCount++
					}
					ctx.Writer().SetPayload(ListResult{
						RequestCount: count,
						AllCount:     listCount,
						PageCount:    PageCount,
						List:         data,
					})
				} else {
					ctx.Writer().SetPayload(err).SetCode(http.StatusNoContent)
				}
			},
		})
	return c
}

func (c *cruder) Get(options ...ICruderOption) (out ICruder) {
	model, _, url, methodUrl := c.modelUrlMethod(http.MethodGet, c.one, c.many, c.apiGet, options...)
	c.api = append(c.api,
		ApiData{
			Url:    url,
			Method: methodUrl,
			Func: func(ctx IContext) {
				id, errInt := ctx.Reader().Param().Int("id")
				if errInt != nil {
					ctx.Writer().SetPayload("not found id query param").SetCode(http.StatusBadRequest)
					return
				}

				var db = c.db
				if unscoped, errUnscoped := ctx.Reader().Param().Bool("unscoped"); errUnscoped == nil && unscoped {
					db = db.Unscoped()
				}
				var (
					count int64
					data  = reflect.New(model).Interface()
				)
				if err := db.Where("id = ?", id).Find(data).Count(&count).Error(); err != nil {
					ctx.Writer().SetPayload(err).SetCode(http.StatusNotFound)
				} else if count == 0 {
					ctx.Writer().SetPayload("no payload")
				} else {
					ctx.Writer().SetPayload(data)
				}
			},
		})
	return c
}

func (c *cruder) Create(options ...ICruderOption) (out ICruder) {
	model, _, url, methodUrl := c.modelUrlMethod(http.MethodPost, c.one, c.many, c.apiCreate, options...)
	c.api = append(c.api,
		ApiData{
			Url:    url,
			Method: methodUrl,
			Func: func(ctx IContext) {
				data := reflect.New(model).Interface()
				if err := ctx.Reader().Body().Json(data); err != nil {
					ctx.Writer().SetPayload(err).SetCode(http.StatusNotAcceptable)
					return
				}

				if err := c.db.Create(data).Error(); err != nil {
					ctx.Writer().SetPayload(err).SetCode(http.StatusBadRequest)
				} else {
					ctx.Writer().SetPayload(data)
				}
			},
		})
	return c
}

func (c *cruder) Delete(options ...ICruderOption) (out ICruder) {
	model, _, url, methodUrl := c.modelUrlMethod(http.MethodDelete, c.one, c.many, c.apiDelete, options...)
	c.api = append(c.api,
		ApiData{
			Url:    url,
			Method: methodUrl,
			Func: func(ctx IContext) {
				id, errInt := ctx.Reader().Param().Int("id")
				if errInt != nil {
					ctx.Writer().SetPayload("not found id query param").SetCode(http.StatusBadRequest)
					return
				}

				var db = c.db
				if unscoped, errUnscoped := ctx.Reader().Param().Bool("unscoped"); errUnscoped == nil && unscoped {
					db = db.Unscoped()
				}

				if err := db.Delete(reflect.New(model).Interface(), id).Error(); err != nil {
					ctx.Writer().SetPayload(err).SetCode(http.StatusBadRequest)
				} else {
					ctx.Writer().SetPayload("ok")
				}
			},
		})
	return c
}

func (c *cruder) Restore(options ...ICruderOption) (out ICruder) {
	model, _, url, methodUrl := c.modelUrlMethod(http.MethodPatch, c.one, c.many, c.apiRestore, options...)
	c.api = append(c.api,
		ApiData{
			Url:    url,
			Method: methodUrl,
			Func: func(ctx IContext) {
				id, errInt := ctx.Reader().Param().Int("id")
				if errInt != nil {
					ctx.Writer().SetPayload("not found id query param").SetCode(http.StatusBadRequest)
					return
				}
				if err := c.db.Unscoped().Model(reflect.New(model).Interface()).Where("id = ?", id).
					Update("deleted_at", nil).Error(); err != nil {
					ctx.Writer().SetPayload(err).SetCode(http.StatusBadRequest)
				} else {
					ctx.Writer().SetPayload("ok")
				}
			},
		})
	return c
}

func (c *cruder) Update(options ...ICruderOption) (out ICruder) {
	model, _, url, methodUrl := c.modelUrlMethod(http.MethodPut, c.one, c.many, c.apiChange, options...)
	c.api = append(c.api,
		ApiData{
			Url:    url,
			Method: methodUrl,
			Func: func(ctx IContext) {
				data := reflect.New(model).Interface()
				if err := ctx.Reader().Body().Json(data); err != nil {
					ctx.Writer().SetPayload(err).SetCode(http.StatusNotAcceptable)
					return
				}
				var db = c.db
				if unscoped, errUnscoped := ctx.Reader().Param().Bool("unscoped"); errUnscoped == nil && unscoped {
					db = db.Unscoped()
				}
				if err := db.Save(data).Error(); err != nil {
					ctx.Writer().SetPayload(err).SetCode(http.StatusBadRequest)
				} else {
					ctx.Writer().SetPayload(data)
				}
			},
		})
	return c
}

func New(model interface{}, db IDB) (out ICruder) {
	tableName := db.TableName(reflect.New(reflect.TypeOf(model)).Interface())
	return &cruder{
		one:        reflect.TypeOf(model),
		many:       reflect.SliceOf(reflect.TypeOf(model)),
		db:         db,
		tableName:  tableName,
		apiList:    "/api/" + strings.ReplaceAll(tableName, ".", "__") + "/list",
		apiGet:     "/api/" + strings.ReplaceAll(tableName, ".", "__") + "/get_by_id",
		apiCreate:  "/api/" + strings.ReplaceAll(tableName, ".", "__") + "/create",
		apiDelete:  "/api/" + strings.ReplaceAll(tableName, ".", "__") + "/delete",
		apiRestore: "/api/" + strings.ReplaceAll(tableName, ".", "__") + "/restore",
		apiChange:  "/api/" + strings.ReplaceAll(tableName, ".", "__") + "/change",
	}
}
