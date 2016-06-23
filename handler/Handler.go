package handler

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
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


func GetUsers(c echo.Context) error {
	db := getDB()
	defer db.Close()

	id := c.Param("id")

	rows, qerr := db.Query("select * from users where id = ?", id)

	defer rows.Close()
	if qerr != nil {
		return c.String(http.StatusInternalServerError, "error")
	}

	for rows.Next() {
		var id string
		if berr := rows.Scan(&id); berr != nil {
			log.Fatal("scan erro: %v", berr)
			return c.String(http.StatusInternalServerError, "error")
		}

		return c.String(http.StatusNotFound, id)
	}
	return c.String(http.StatusNotFound, "user not found")
}
