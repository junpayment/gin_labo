package models

import (
  "github.com/go-xorm/xorm"
  _ "github.com/mattn/go-sqlite3"
  "log"
)

const (
  dbName = "gin_labo.sqlite3"
)

var engine *xorm.Engine

// init engine
func Init() {
  var err error
  engine, err = xorm.NewEngine("sqlite3", dbName)
  if err != nil {
    log.Fatal(err)
  }

  if err = engine.DB().Ping(); err != nil {
    log.Fatal(err)
  }
}

// close engine
func Close() {
  engine.Clone()
}
