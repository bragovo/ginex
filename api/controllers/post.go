package controllers

import (
  "ginex/api/models"

  "github.com/gin-gonic/gin"
  "net/http"
)

func GetPosts(c *gin.Context) {
  var posts []models.Post
  models.GetPosts(&posts)

  c.JSON(http.StatusOK, gin.H{
    "posts": posts,
  })
}

func GetPost(c *gin.Context) {
  id := c.Param("id")
  var post models.Post
  models.GetPost(&post, id)

  c.JSON(http.StatusOK, gin.H{
    "post": post,
  })
}
