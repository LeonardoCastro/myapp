package controllers

import (
      "github.com/jmoiron/sqlx"
      _ "github.com/lib/pq"
       "log"
      // "database/sql"
      // "golang.org/x/crypto/bcrypt"
      "github.com/revel/revel"
    	//"github.com/revel/modules/db/app"
)

 var db *sqlx.DB

 // var schema = `
 // CREATE TABLE Users(
 //   Id key int,
 //   Username text,
 //   Password text,
 //   HashedPassword bytea
 //   )`

type SqlxController struct {
	*revel.Controller
  Txn *sqlx.Tx
}


type User struct {
  Id int
  Username string
  Password string
}


func InitDB() {
  var err error
  db, err = sqlx.Open("postgres", "user=Leo dbname=testdb sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }

  // db.MustExec(schema)
}


func (c *SqlxController) Begin() revel.Result {
  txn, err := db.Beginx()
  if err != nil {
    panic(err)
  }
  c.Txn = txn
  return nil
}
