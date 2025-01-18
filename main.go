package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gsh519/echo-go-todo/controllers"
	"github.com/gsh519/echo-go-todo/services"
	"github.com/labstack/echo/v4"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_DATABASE")
	dbConn     = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	ser := services.NewAppService(db)
	con := controllers.NewTodoController(ser)

	e := echo.New()

	e.GET("/todo", con.FetchTodoHandler())
	e.POST("/todo", con.CreateTodoHandler())
	e.PUT("/todo/:id", con.UpdateTodoHandler())
	e.PUT("/todo/:id/done", con.DoneTodoHandler())
	e.DELETE("/todo/:id", con.DeleteTotoHandler())
	e.GET("/todo/done", con.FetchDoneTodoHandler())

	e.Logger.Fatal(e.Start(":8888"))
}
