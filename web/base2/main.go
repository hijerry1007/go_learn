package main

import (
	"ivy"
	"net/http"
)

func main() {
	router := ivy.New()
	router.Get("/", func(c *ivy.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	router.Run(":9999")
}
