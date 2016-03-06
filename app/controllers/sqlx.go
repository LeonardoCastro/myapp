package controllers

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Discrete use
	"log"
	// "database/sql"
	// "golang.org/x/crypto/bcrypt"
	"github.com/revel/revel"
	//"github.com/revel/modules/db/app"
)

// db: database
var db *sqlx.DB

// var schema = `
// CREATE TABLE Users(
//   Id key int,
//   Username text,
//   Password text,
//   HashedPassword bytea
//   )`

// SqlxController struct
type SqlxController struct {
	*revel.Controller
	Txn *sqlx.Tx
}

// User type used to link to the database
type User struct {
	ID       int
	Username string
	Password string
}

// InitDB connects to the database
func InitDB() {
	var err error
	db, err = sqlx.Open("postgres", "user=Leo dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// db.MustExec(schema)
}

// Begin the database with a transaction
func (c *SqlxController) Begin() revel.Result {
	txn, err := db.Beginx()
	//db.MustExec("set transaction isolation level read uncommitted")
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

// Commit a transaction
func (c *SqlxController) Commit(Txn *sqlx.Tx) revel.Result {
	Txn.MustExec("commit transaction")
	return nil
}

// RollBack a transaction
func (c *SqlxController) RollBack(Txn *sqlx.Tx) revel.Result {
	Txn.MustExec("rollback transaction")
	return nil
}
