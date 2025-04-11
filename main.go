package main

import (
	"nexmedis_project/controller"
	"nexmedis_project/middleware"
	"nexmedis_project/model"
	"nexmedis_project/repository"
	"nexmedis_project/service"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	db := model.BuildDB()
	defer func() {
		if sqlDB, err := db.DB(); err != nil {
			panic(err)
		} else {
			_ = sqlDB.Close()
		}
	}()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authController := controller.NewAuthController(userService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	router.POST("/register", authController.Register)

	router.POST("/login", authController.Login)

	protected := router.Group("/user")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/my", authController.CurrentUser)

	protected.GET("/products", productController.GetAllProducts)
	protected.GET("/products/:id", productController.GetProductByID)
	protected.GET("/products/search", productController.SearchProducts)

	router.Run(":8080")

}
