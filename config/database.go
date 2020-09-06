package config


import (
  "github.com/go-pg/pg/v10"
  "github.com/go-pg/pgext"
)

var DB *pg.DB

func init(){
  DB = pg.Connect(&pg.Options{
    Addr: ":5432",
    User: "brg",
    Password: "",
    Database: "go_dev",
  })

  DB.AddQueryHook(&pgext.DebugHook{
    Verbose: true,
  })
}
