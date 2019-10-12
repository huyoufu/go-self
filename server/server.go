package server

import (
	"fmt"
	"github.com/huyoufu/go-self/logger"
	"github.com/huyoufu/go-self/router"
	"github.com/huyoufu/go-self/session"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	port           int
	serverName     string
	cors           bool
	sessionManager *session.Manager
	session        bool
	pl             *router.Pipeline
}

func DefaultServer() *Server {
	server := NewServer()
	//添加日志拦截
	server.AppendValveF(func(ctx router.Context) bool {
		fmt.Println(time.Now().Format("009/01/23 01:23:23") + "|" + ctx.ClientIP() + "|" + ctx.Req().RequestURI)
		return true
	})
	return server
}
func (s *Server) FrontValve(valves ...router.Valve) {
	for _, v := range valves {
		s.pl.First(v)
	}
}
func (s *Server) AppendValve(valves ...router.Valve) {
	for _, v := range valves {
		s.pl.Last(v)
	}
}
func (s *Server) FrontValveF(vfs ...func(ctx router.Context) bool) {
	for _, v := range vfs {
		s.pl.FirstPF(v)
	}
}
func (s *Server) AppendValveF(vfs ...func(ctx router.Context) bool) {
	for _, v := range vfs {
		s.pl.LastPF(v)
	}
}
func NewServer() *Server {
	return &Server{
		port:       8847,
		serverName: "jk1123.com",
		cors:       false,
		pl:         router.New(),
	}
}

func (s *Server) EnableCors() {
	s.cors = true
}
func (s *Server) EnableSession() {
	s.sessionManager = session.DefaultManager()
	s.session = true

}
func (s *Server) Port(port int) {
	s.port = port
}
func (s *Server) Start() {
	logger.InfoF("server will start on port: %d", s.port)
	if s.session {
		s.sessionManager.StartGC()
	}
	http.Handle("/", s)
	e := http.ListenAndServe(":"+strconv.Itoa(s.port), nil)
	if e != nil {
		panic(e)
	}

}

func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("server", s.serverName)
	if s.cors {
		ori := req.Header.Get("origin")
		//fmt.Println(ori)
		//允许跨域访问
		resp.Header().Set("Access-Control-Allow-Origin", ori)
		//允许cookie跨域
		resp.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	vr := router.Dispatcher[req.Method]
	_, h, params := vr.Tree.Search(req.URL.Path)
	//fmt.Println("hahha")
	if h == nil {
		_, h, params = router.Dispatcher["ANY"].Tree.Search(req.URL.Path)
	}
	if h == nil {
		http.NotFoundHandler().ServeHTTP(resp, req)
		return
	}
	httpContext := router.NewHttpContext(req, resp, params, nil)
	if s.session {
		//支持session
		httpContext.Session = initSession(s.sessionManager, req, resp)
	}

	s.pl.Invoke(httpContext)
	if httpContext.IsEnd() {
		httpContext.WriteString("非法请求")
		return
	}
	h.Service(httpContext)

}

func initSession(manager *session.Manager, request *http.Request, resp http.ResponseWriter) (s session.Session) {
	cookies := request.Cookies()
	var sessionCookie *http.Cookie = nil

	if cookies != nil && len(cookies) > 0 {
		for _, c := range cookies {
			if session.SessionCookieName == c.Name {
				sessionCookie = c
				break
			}
		}
	}
	if sessionCookie != nil {
		s = manager.GetSession(sessionCookie.Value)
	} else {
		s = manager.NewSession()
		sessionCookie = &http.Cookie{
			Name:  session.SessionCookieName,
			Value: s.Id(),
			Path:  "/",
		}
		resp.Header().Set("set-cookie", sessionCookie.String())
	}
	return
}
