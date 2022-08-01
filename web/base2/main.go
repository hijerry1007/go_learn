package main

import (
	"fmt"
	"html/template"
	"ivy"
	"net/http"
	"time"
)

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := ivy.Default()
	r.SetFunMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("views/*")
	r.Static("/assets", "./static")

	r.GET("/panic", func(c *ivy.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	r.GET("/index", func(c *ivy.Context) {
		c.HTML(http.StatusOK, "user.html", ivy.H{
			"title": "i am the title",
		})
	})
	v1 := r.Group("/v1")
	{
		v1Child := v1.Group("/v1Child")

		v1Child.GET("/", func(c *ivy.Context) {
			c.HTML(http.StatusOK, "user.html", ivy.H{
				"title": "i am the title",
			})
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
