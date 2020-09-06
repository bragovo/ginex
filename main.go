package main

import (
  "ginex/api/controllers"

  "github.com/gin-gonic/gin"
  "net/http"
)

func main() {
  router := gin.Default()
  router.Static("/assets", "./assets")
  router.LoadHTMLGlob("templates/*")

  api := router.Group("/api")
  {
    api.GET("posts", controllers.GetPosts)
    api.GET("posts/:id", controllers.GetPost)
  }

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Main website",
    })
  })

  router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
