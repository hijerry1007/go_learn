package main

import (
	"ivy"
	"net/http"
)

func main() {
	r := ivy.New()
	r.GET("/index", func(c *ivy.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1Child := v1.Group("/v1Child")

		v1Child.GET("/", func(c *ivy.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1Child.GET("/hello", func(c *ivy.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})

	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *ivy.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *ivy.Context) {
			c.JSON(http.StatusOK, ivy.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
