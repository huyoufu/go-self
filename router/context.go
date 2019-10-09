package router

import (
	"context"
	_ "github.com/huyoufu/go-self/json"
	"github.com/huyoufu/go-self/session"
	"github.com/json-iterator/go"
	"net/http"
	"net/url"
)

type Context interface {
	Req() *http.Request
	Resp() http.ResponseWriter
	PathParamValue(name string) string
	Bind(bean interface{}) error
	WriteString(string)
	WriteJson(bean interface{})
	SessionGet(key string) interface{}
	SessionSet(key string, value interface{})
	SessionInvalidate()
}
type PathParam map[string]string
type HttpContext struct {
	req     *http.Request
	resp    http.ResponseWriter
	params  PathParam
	Session session.Session
	context context.Context
}

func NewHttpContext(req *http.Request, resp http.ResponseWriter, params PathParam, Session session.Session) *HttpContext {
	return &HttpContext{
		req,
		resp,
		params,
		nil,
		nil,
	}
}

func (hctx *HttpContext) SessionGet(key string) interface{} {
	if hctx.Session != nil {
		return hctx.Session.Get(key)
	}
	return nil
}

func (hctx *HttpContext) SessionSet(key string, value interface{}) {
	if hctx.Session != nil {
		hctx.Session.Set(key, value)
	}
}

func (hctx *HttpContext) SessionInvalidate() {
	hctx.Session.Invalidate()
}

func (hctx *HttpContext) PathParamValue(name string) string {
	return hctx.params[name]
}
func (hctx *HttpContext) Req() *http.Request {
	return hctx.req
}
func (hctx *HttpContext) Bind(bean interface{}) error {
	hctx.req.ParseForm()
	//e := bind.Bind(bean, hctx.req.PostForm)
	e := Bind(bean, hctx.req.Form)
	return e
}
func copyValues(dst, src url.Values) {
	for k, vs := range src {
		for _, value := range vs {
			dst.Add(k, value)
		}
	}
}
func (hctx *HttpContext) Resp() http.ResponseWriter {
	return hctx.resp
}
func (hctx *HttpContext) WriteString(s string) {
	hctx.resp.Write([]byte(s))
}
func (hctx *HttpContext) WriteJson(bean interface{}) {

	bytes, _ := jsoniter.Marshal(bean)
	hctx.resp.Write(bytes)
}
