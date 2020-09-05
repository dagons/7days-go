package gee

import (
	"net/http"
)

//框架引擎主体
type Engine struct {
	route *Route
}

func New() *Engine {
	return &Engine{route: newRoute()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.route.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//实现ServeHTTP方法 满足Handler接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.route.handle(c)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
