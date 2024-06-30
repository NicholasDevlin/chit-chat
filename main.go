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

	customerRepository := repositories.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)
	e.POST("/customer", customerController.CreateCustomer)
	e.GET("/customer", customerController.GetAllCustomer)

	productRepository := repositories.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)
	e.POST("/product", productController.CreateProduct)
	e.GET("/product", productController.GetAllProduct)

	userRepository := repositories.NewUsersRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	e.POST("/user/register", userController.RegisterUsers)
	e.POST("/user/login", userController.LoginUser)
	e.GET("/user", userController.GetAllUser)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.APP_PORT)))
}
