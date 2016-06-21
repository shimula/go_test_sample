package main

import (
	"net/http"
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"fmt"
)

func getDB() (db *sql.DB) {

	dbHost := os.Getenv("MYSQL_HOST")
	if dbHost == "" {
		dbHost = "192.168.99.100"
	}

	dbUser := os.Getenv("MYSQL_USER")
	if dbUser == "" {
		dbUser = "testuser"
	}
	dbPort := os.Getenv("MYSQL_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	if dbPassword == "" {
		dbPassword = "password"
	}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/test", dbUser, dbPassword, dbHost, dbPort))
	if err != nil {
		log.Fatal("open erro: %v", err)
	}
	return db
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "hello, world")
}

func getUsers(c echo.Context) error {
	db := getDB()
	defer db.Close()

	id := c.Param("id")

	rows, qerr := db.Query("select * from users where id = ?", id)

	defer rows.Close()
	if qerr != nil {
		return qerr
	}

	results :=  ""
	for rows.Next() {
		var id string
		if berr := rows.Scan(&id); berr != nil {
			log.Fatal("scan erro: %v", berr)
		}

		results = results + id + "\n"
	}

	if results != "" {
		return c.String(http.StatusOK, results)
	}
	return echo.ErrNotFound
}

func main() {
	e := echo.New()
	e.GET("/", index)
	e.GET("/users/:id", getUsers)
	e.Run(standard.New(":8888"))
}
