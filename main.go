package main

import (
	"Assignment/app"
	"Assignment/handlers"
	"Assignment/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middleware.PanicHandling())

	db := app.Connection()
	app.AutoMigration(db)

	productAPIs := router.Group("/product")
	productAPIs.POST("/add", handlers.AddProductHandler)
	productAPIs.GET("/list-all", handlers.ReadProductsHandler)
	productAPIs.GET("/:id", handlers.ReadProductHandler)

	orderAPIs := router.Group("/order")
	orderAPIs.POST("/add", handlers.PlaceOrderHandler)
	orderAPIs.GET("/:id", handlers.ReadOrderHandler)
	orderAPIs.PATCH("/status/:id", handlers.UpdateStatusHandler)
	orderAPIs.PATCH("/cancel/:id", handlers.CancelOrderHandler)
	router.Run()
}
