package main

import (
	"github.com/gin-gonic/gin"
)

// GinEngine returns an instance of the gin Engine.
func GinEngine() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("ui/*")

	r.Static("a", "a")

	r.GET("/", func(c *gin.Context) { c.HTML(200, "index.html", nil) })

	r.GET("/k/:k", func(c *gin.Context) { c.HTML(200, "list.html", gin.H{"klasse": c.Param("k")}) })

	api := r.Group("api")
	{
		api.GET("/", apiRoot)

		api.GET("/:klasse", apiParser)
	}

	return r
}

func main() {
	GinEngine().Run(":5000")
}
