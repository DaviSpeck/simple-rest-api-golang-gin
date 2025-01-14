package main

import (
	"simple-rest-api-golang-gin/controller"
	"simple-rest-api-golang-gin/db"
	"simple-rest-api-golang-gin/repository"
	"simple-rest-api-golang-gin/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Repositories
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Usecases
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// Controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)

	server.Run(":8000")

}
