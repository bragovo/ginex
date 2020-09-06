package controllers

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func GetPosts(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "POSTS",
  })
}

func GetPost(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "POST",
  })
}
