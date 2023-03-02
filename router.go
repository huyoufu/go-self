package self

import (
	"github.com/huyoufu/go-self/rtree"
	"strings"
)

type Routers map[string]router

var Dispatcher = Routers{
	"GET":     router{"get请求", rtree.NewRoot()},
	"POST":    router{"post请求", rtree.NewRoot()},
	"PUT":     router{"put请求", rtree.NewRoot()},
	"DELETE":  router{"delete请求", rtree.NewRoot()},
	"ANY":     router{"any请求", rtree.NewRoot()},
	"OPTIONS": router{"options请求", rtree.NewRoot()},
}

type Group struct {
	name, gpath string
	pl          *Pipeline
}
type router struct {
	name string
	Tree *rtree.Node
}

func NewGroup(name, groupPath string) *Group {
	return &Group{
		name,
		groupPath,
		NewPipeline(),
	}
}

func (rg *Group) AppendValve(valves ...Valve) {
	for _, v := range valves {
		rg.pl.Last(v)
	}
}

func (rg *Group) AppendValveF(vfs ...func(ctx Context)) {
	for _, v := range vfs {
		rg.pl = rg.pl.LastPF(v)
	}
}

func (rg *Group) Get(path string, h func(ctx Context)) {
	AddRouterHandFuncWithPipeline("GET", rg.gpath+path, h, rg.pl)
}
func (rg *Group) Post(path string, h func(ctx Context)) {
	AddRouterHandFuncWithPipeline("POST", rg.gpath+path, h, rg.pl)
}
func (rg *Group) Put(path string, h func(ctx Context)) {
	AddRouterHandFuncWithPipeline("PUT", rg.gpath+path, h, rg.pl)
}
func (rg *Group) Delete(path string, h func(ctx Context)) {
	AddRouterHandFuncWithPipeline("DELETE", rg.gpath+path, h, rg.pl)
}
func (rg *Group) Any(path string, h func(ctx Context)) {
	AddRouterHandFuncWithPipeline("ANY", rg.gpath+path, h, rg.pl)
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
func AddRouterHandFuncWithPipeline(method, pattern string, h HandlerFunc, pl *Pipeline) {
	p := Cleanup(pattern)
	r := Dispatcher[method]

	handlerPipeline := HandlerPipeline{h, pl}
	r.Tree.AddNode(p, &handlerPipeline)
}
func AddRouterHandFunc(method, pattern string, h HandlerFunc) {
	AddRouterHandFuncWithPipeline(method, pattern, h, NewPipeline())
}

func init() {
	AddRouterHandFunc("GET", "/hello", func(ctx Context) {
		//获取当前服务器的相关信息

		//pprof.StartCPUProfile(ctx.Resp())
		//pprof.StopCPUProfile()
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

func prefix(s string, prefix string) string {
	if !strings.HasPrefix(s, prefix) {
		return prefix + s
	}
	return s
}

func suffix(s string, suffix string) string {
	if !strings.HasSuffix(s, suffix) {
		return s + suffix
	}
	return s
}
