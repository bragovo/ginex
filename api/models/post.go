package models

import (
  "ginex/config"
)

type Post struct {
  Id int64
  Title string
}

func GetPosts(posts *[]Post) {
  config.DB.Model(posts).Select()
}

func GetPost(post *Post, id string) {
  config.DB.Model(post).Where("id = ?", id).Select()
}
