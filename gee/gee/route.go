package gee

import (
	"log"
	"net/http"
)

//路由指向的方法
type HandlerFunc func(*Context)
type Route struct {
	handlers map[string]HandlerFunc
}

func newRoute() *Route {
	return &Route{handlers: make(map[string]HandlerFunc)}
}
func (r *Route) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "_" + pattern
	r.handlers[key] = handler
}

func (r *Route) handle(c *Context) {
	key := c.Method + "_" + c.Path
	if handler, err := r.handlers[key]; err {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "40 not FOUND:%s\n", c.Path)
	}
}
