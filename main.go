package main

import (
	"fmt"
	"myapp/backend/config"
	"myapp/backend/controller"
	"myapp/backend/repositories"
	"myapp/backend/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	appConfig, dbConfig := config.InitConfig()
	db := config.StartDB(dbConfig)
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	repository := repositories.NewCustomerRepository(db)
	service := service.NewCustomerService(repository)
	controller := controller.NewCustomerController(service)
	e.POST("/customer", controller.CreateCustomer)
	e.GET("/customer", controller.GetAllCustomer)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.APP_PORT)))
}