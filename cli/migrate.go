package main

import (
  "os"
  "log"
  "github.com/urfave/cli"
  "github.com/go-xorm/xorm"
  _ "github.com/mattn/go-sqlite3"
  "github.com/go-xorm/xorm/migrate"
  "github.com/junpayment/gin_labo/apps/models"
)

const (
  dbName = "gin_labo.sqlite3"
)

var (
  migrations = []*migrate.Migration{
    {
      ID: "201712310302",
      Migrate: func(tx *xorm.Engine) error {
        return tx.Sync2(&models.User{})
      },
      Rollback: func(tx *xorm.Engine) error {
        return tx.DropTables(&models.User{})
      },
    },
  }
)

func main() {
  app := cli.NewApp()
  app.Name = "migrate"
  app.Usage = "migrate all schemas"
  app.Action = execMigrate
  app.Run(os.Args)
}

// migration
func execMigrate(c *cli.Context) error {
  db, err := xorm.NewEngine("sqlite3", dbName)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  if err = db.DB().Ping(); err != nil {
    log.Fatal(err)
  }

  m := migrate.New(db, migrate.DefaultOptions, migrations)

  err = m.Migrate()

  return nil
}
