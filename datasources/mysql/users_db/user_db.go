package users_db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client   *sql.DB
	err      error
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "root", "127.0.0.1:3306", "users_db")

	fmt.Println("Setting up data source")
	fmt.Println(dataSourceName)
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database succesfully configured")
}
