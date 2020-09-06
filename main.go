package main

import (
  "ginex/api/controllers"

  "github.com/gin-gonic/gin"
  // "github.com/go-pg/pg"
)

func main() {
  router := gin.Default()

  api := router.Group("/api")
  {
    api.GET("posts", controllers.GetPosts)
    api.GET("posts/:id", controllers.GetPost)
  }

  router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
