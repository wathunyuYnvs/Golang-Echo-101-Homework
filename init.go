package myecho

import (
	"myecho/db"
	"myecho/middlewares"
	"myecho/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s"
)

// @title Golang - Echo 101
// @description User Management API
// @version 1.0
// @host localhost:1323
// @BasePath /
func InitApp() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	resource, err := db.Init()
	if err != nil {
		logrus.Error(err)
	}

	defer resource.Close()
	prRoute := e.Group("/private")
	prRoute.Use(middleware.JWTWithConfig(middlewares.Config()))
	user.PrivateEndpointHadlerInit(prRoute, resource)

	pbRoute := e.Group("/public")
	user.PublicEndpointHadlerInit(pbRoute, resource)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
