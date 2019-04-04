package router

import (
	"strings"
)

type Routers map[string]router

var Dispatcher = Routers{
	"GET":     router{"get请求", NewRoot()},
	"POST":    router{"post请求", NewRoot()},
	"PUT":     router{"put请求", NewRoot()},
	"DELETE":  router{"delete请求", NewRoot()},
	"ANY":     router{"any请求", NewRoot()},
	"OPTIONS": router{"options请求", NewRoot()},
}

type RouterGroup struct {
	name, gpath string
}
type router struct {
	name string
	Tree *Node
}

func NewGroup(name, groupPath string) *RouterGroup {
	return &RouterGroup{
		name,
		groupPath,
	}
}

func (rg *RouterGroup) Get(path string, h func(ctx Context)) {
	AddRouterHandFunc("GET", rg.gpath+path, h)
}
func (rg *RouterGroup) Post(path string, h func(ctx Context)) {
	AddRouterHandFunc("POST", rg.gpath+path, h)
}
func (rg *RouterGroup) Put(path string, h func(ctx Context)) {
	AddRouterHandFunc("PUT", rg.gpath+path, h)
}
func (rg *RouterGroup) Delete(path string, h func(ctx Context)) {
	AddRouterHandFunc("DELETE", rg.gpath+path, h)
}
func (rg *RouterGroup) Any(path string, h func(ctx Context)) {
	AddRouterHandFunc("ANY", rg.gpath+path, h)
}

func Get(pattern string, h func(ctx Context)) {
	AddRouterHandFunc("GET", pattern, h)
}
func Post(pattern string, h func(ctx Context)) {
	AddRouterHandFunc("POST", pattern, h)
}
func Put(pattern string, h func(ctx Context)) {
	AddRouterHandFunc("PUT", pattern, h)
}
func Delete(pattern string, h func(ctx Context)) {
	AddRouterHandFunc("DELETE", pattern, h)
}
func Any(pattern string, h func(ctx Context)) {
	AddRouterHandFunc("ANY", pattern, h)
}
func Options(pattern string, h func(ctx Context)) {
	AddRouterHandFunc("OPTIONS", pattern, h)
}

func AddRouterHandFunc(method, pattern string, h func(ctx Context)) {
	p := Cleanup(pattern)
	r := Dispatcher[method]
	r.Tree.addNode(p, HandlerFunc(h))
}

func init() {
	AddRouterHandFunc("GET", "/hello", func(ctx Context) {

		ctx.Resp().Write([]byte("<div style='text-align:center'><a href='http://www.jk1123.com'>www.jk1123.com 怪蜀黍</a></div>"))
	})

}
func Cleanup(pattern string) string {
	//if null  set  "/"
	if pattern == "" || pattern == "." {
		return "/"
	}
	pattern = strings.Replace(pattern, "\\", "/", -1)
	pattern = strings.Replace(pattern, "//", "/", -1)

	//if start not "/" then add "/"
	if pattern[0] != '/' {
		pattern = "/" + pattern
	}
	//trim right "/"
	pattern = strings.TrimRight(pattern, "/")
	//fmt.Println(pattern)
	//clean right "/"  if before "/" we must set pattern="/"
	pattern = prefix(pattern, "/")
	return pattern
}