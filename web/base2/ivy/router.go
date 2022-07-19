package ivy

import (
	"net/http"
)

type router struct {
	handlers map[string]Handler
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]Handler),
	}
}

func (r *router) addRoute(method string, pattern string, handler Handler) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
