package self

import (
	"context"
	_ "github.com/huyoufu/go-self/json"
	"github.com/huyoufu/go-self/resolver"
	"github.com/huyoufu/go-self/session"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Context interface {
	Req() *http.Request
	Resp() http.ResponseWriter
	PathParamValue(name string) string
	Bind(bean interface{}) error
	GetParameter(name string) string
	GetParameterValues(name string) []string
	WriteString(string)
	WriteJson(bean interface{})
	SessionGet(key string) interface{}
	SessionSet(key string, value interface{})
	SessionInvalidate()
	Abort()
	IsEnd() bool
	ClientIP() string
	Next() //执行下一个拦截器

}
type PathParam map[string]string
type HttpContext struct {
	req     *http.Request
	resp    http.ResponseWriter
	params  PathParam
	Session session.Session
	context context.Context
	end     bool
	hp      *HandlerPipeline
}

func (hctx *HttpContext) Next() {
	hctx.hp.Invoke(hctx)
}

func NewHttpContext(req *http.Request, resp http.ResponseWriter, params PathParam, hp *HandlerPipeline) *HttpContext {
	return &HttpContext{
		req,
		resp,
		params,
		nil,
		nil,
		false,
		hp,
	}
}

//此方法不保证 完全正确 因为有时候浏览器发送ajax请求的时候 有可能不会带该请求头
//此方法只作为临时性实现
func (hctx *HttpContext) IsAjax() bool {
	xrw := hctx.req.Header.Get("X-Requested-With")
	if xrw != "" && "XMLHttpRequest" == xrw {
		return true
	}
	return false
}

func (hctx *HttpContext) GetParameter(name string) string {
	hctx.req.ParseForm()
	return hctx.req.Form.Get(name)
}
func (hctx *HttpContext) GetParameterValues(name string) []string {
	if hctx.req.Method == http.MethodGet {
		//get请求 直接去获取查询字符串
		vs := hctx.req.URL.Query()
		for k, v := range vs {
			if k == name {
				return v
			}
		}
	} else {
		hctx.req.ParseForm()
		vs := hctx.req.Form
		for k, v := range vs {
			if k == name {
				return v
			}
		}
	}
	return nil
}

func (hctx *HttpContext) ClientIP() string {
	return hctx.req.RemoteAddr
}

func (hctx *HttpContext) IsEnd() bool {
	return hctx.end
}

func (hctx *HttpContext) Abort() {

	hctx.end = true
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
	contentType := hctx.Req().Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		bytes, _ := ioutil.ReadAll(hctx.Req().Body)
		return jsoniter.Unmarshal(bytes, bean)
	} else {
		hctx.req.ParseForm()
		return resolver.Bind(bean, hctx.req.Form)
	}

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
