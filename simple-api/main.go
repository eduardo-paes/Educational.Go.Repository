package main

import (
	"github.com/eduardo-paes/clean-go/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Main function
func main() {
	// Echo instance
	e:= echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	uc := controllers.NewUserController()

	// Routes
	e.POST("/users", uc.CreateUser)
	e.GET("/users/:id", uc.GetUser)
	e.GET("/users", uc.GetUsers)
	e.PUT("/users/:id", uc.UpdateUser)
	e.DELETE("/users/:id", uc.DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}