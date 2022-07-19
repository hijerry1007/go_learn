package ivy

import (
	"net/http"
)

type Handler func(c *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) Get(pattern string, handler Handler) {
	e.router.addRoute("GET", pattern, handler)
}

func (e *Engine) Post(pattern string, handler Handler) {
	e.router.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)

	e.router.handle(c)

}
