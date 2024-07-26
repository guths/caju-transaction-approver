package configs

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type dbConfig struct {
	host     string
	name     string
	user     string
	password string
	db       string
}

func (db *dbConfig) validate() {
	if db.user == "" {
		panic(fmt.Errorf("user is empty in db %s", db.name))
	}

	if db.password == "" {
		panic(fmt.Errorf("password is empty in db %s", db.name))
	}

	if db.db == "" {
		panic(fmt.Errorf("database name is empty in db %s", db.name))
	}
}

func init() {
	db := dbConfig{
		user:     os.Getenv("MYSQL_USER"),
		password: os.Getenv("MYSQL_PASSWORD"),
		db:       os.Getenv("MYSQL_DATABASE"),
		host:     os.Getenv("MYSQL_HOST"),
	}

	db.validate()

	DB = connectDB(buildUrl(db))
}

func buildUrl(db dbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db.user, db.password, db.host, db.db)
}

func connectDB(url string) *sql.DB {
	var err error

	db, err := sql.Open("mysql", url)

	if err != nil {
		panic(err.Error())
	}

	// defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	return db
}
