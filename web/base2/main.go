package main

import (
	"fmt"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.path {
	case "/":
		fmt.Fprintf(w, "path = %q", req.URL.path)
	}
}
func main() {

}
