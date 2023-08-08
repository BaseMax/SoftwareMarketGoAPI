package main

import (
	"Software-Market-Go-API/delivery/httpserver"
	"Software-Market-Go-API/delivery/httpserver/api"
	"Software-Market-Go-API/delivery/httpserver/panel"
	"Software-Market-Go-API/repository/mysql"
	"Software-Market-Go-API/service/softwareservice"
	"Software-Market-Go-API/service/userservice"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Repo *mysql.MySQLDB

func init() {

	Repo = mysql.ConnectDatabase()
}

func main() {

	err := godotenv.Load("local.env")
	if err != nil {
		panic(fmt.Errorf("some error occured. Err: %s", err))
	}

	Repo.Migrate()

	softwareSvc := softwareservice.New(Repo)
	userSvc := userservice.New(Repo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	aHandler := api.NewHandler(*userSvc, *softwareSvc)
	pHandler := panel.NewHandler(*userSvc, *softwareSvc)

	httpserver.Routes(aHandler, pHandler, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
