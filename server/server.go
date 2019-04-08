package server

import (
	"github.com/huyoufu/go-self/logger"
	"github.com/huyoufu/go-self/router"
	"github.com/huyoufu/go-self/session"
	"net/http"
	"strconv"
)

type Server struct {
	port           int
	serverName     string
	cors           bool
	corsWhiteList  []string
	sessionManager *session.Manager
	session        bool
}

func NewServer() *Server {

	return &Server{
		port:       8847,
		serverName: "jk",
		cors:       false,
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
func (s *Server) AddWhiteList(domain string) {

	s.corsWhiteList = append(s.corsWhiteList, domain)
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

	h.Service(httpContext)

}
func isAjax(request *http.Request) bool {
	xrw := request.Header.Get("X-Requested-With")
	if xrw != "" && "XMLHttpRequest" == xrw {
		return true
	}
	return false

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
