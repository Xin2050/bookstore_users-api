package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

const (
	mysql_username = "mysql_username"
	mysql_password = "mysql_password"
	mysql_host     = "mysql_host"
	mysql_db       = "mysql_db"
)

var (
	Client *sql.DB
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv(mysql_username)
	password := os.Getenv(mysql_password)
	host := os.Getenv(mysql_host)
	db := os.Getenv(mysql_db)
	// create connection to server
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, password, host, db)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	Client.SetConnMaxLifetime(time.Minute * 3)
	Client.SetMaxOpenConns(10)
	Client.SetMaxIdleConns(10)

	log.Println("database successfully configured")

}
