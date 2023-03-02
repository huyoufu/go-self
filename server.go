package self

import (
	. "github.com/huyoufu/go-self/logger"
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
	pl             *Pipeline
}

func DefaultServer() *Server {
	server := NewServer()
	server.EnableSession()
	server.EnableCors()
	//添加日志拦截
	server.AppendValveF(func(ctx Context) {

		start := time.Now().UnixNano()

		ctx.Next()
		end := time.Now().UnixNano()
		i := end - start
		Log.Infof("Request cost:%d", i/1000/1000)
	}, func(ctx Context) {
		//logger.Log.Infof("\x1b[0;31m" + ctx.ClientIP() + " | " + ctx.Req().RequestURI + "\x1b[0m")
		Log.Info(ctx.ClientIP() + " | " + ctx.Req().RequestURI)
		ctx.Next()
	})
	return server
}

func (s *Server) AppendValve(valves ...Valve) {
	for _, v := range valves {
		s.pl.Last(v)
	}
}
func (s *Server) AppendValveF(vfs ...ValveFunc) {
	for _, v := range vfs {
		s.pl = s.pl.LastPF(v)
	}
}
func NewServer() *Server {
	return &Server{
		port:       8847,
		serverName: "self",
		cors:       false,
		pl:         NewPipeline(),
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
	Log.Infof("server will start on port: %d", s.port)
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
	var hp *HandlerPipeline
	vr := Dispatcher[req.Method]
	_, _hp, params := vr.Tree.Search(req.URL.Path)

	if _hp == nil {
		_, _hp, params = Dispatcher["ANY"].Tree.Search(req.URL.Path)
		if _hp != nil {
			hp = _hp.(*HandlerPipeline)
		}
	} else {
		hp = _hp.(*HandlerPipeline)
	}
	if hp == nil {
		http.NotFoundHandler().ServeHTTP(resp, req)
		return
	}

	//获取sever的pipeline
	pl := Compose(s.pl, hp.GetPipeLine())
	nhp := NewRouterHandlerPipeline(pl)
	nhp.Handler = hp.Handler

	httpContext := NewHttpContext(req, resp, params, nhp)
	if s.session {
		//支持session
		httpContext.Session = initSession(s.sessionManager, req, resp)
	}

	//开始执行整个链
	nhp.Invoke(httpContext)

}

func initSession(manager *session.Manager, request *http.Request, resp http.ResponseWriter) (s session.Session) {
	cookies := request.Cookies()
	var sessionCookie *http.Cookie = nil

	if cookies != nil && len(cookies) > 0 {
		for _, c := range cookies {
			if session.CookieName == c.Name {
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
			Name:  session.CookieName,
			Value: s.Id(),
			Path:  "/",
		}
		resp.Header().Set("set-cookie", sessionCookie.String())
	}
	return
}
