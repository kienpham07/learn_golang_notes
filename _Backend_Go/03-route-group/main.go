package main

import (
	v1handler "03-route-group/internal/api/v1/handler"
	v2handler "03-route-group/internal/api/v2/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userHandlerV1 := v1handler.NewUserHandler()
	r.GET("/api/v1/users", userHandlerV1.GetUsersV1)
	r.GET("/api/v1/users/:id", userHandlerV1.GetUsersByIdV1)
	r.POST("/api/v1/users", userHandlerV1.PostUsersV1)
	r.PUT("/api/v1/users/:id", userHandlerV1.PutUsersV1)
	r.DELETE("/api/v1/users/:id", userHandlerV1.DeleteUsersV1)

	productHandlerV1 := v1handler.NewProductHandler()
	r.GET("/api/v1/products", productHandlerV1.GetProductsV1)
	r.GET("/api/v1/products/:id", productHandlerV1.GetProductsByIdV1)
	r.POST("/api/v1/products", productHandlerV1.PostProductsV1)
	r.PUT("/api/v1/products/:id", productHandlerV1.PutProductsV1)
	r.DELETE("/api/v1/products/:id", productHandlerV1.DeleteProductsV1)

	userHandlerV2 := v2handler.NewUserHandler()
	r.GET("/api/v2/users", userHandlerV2.GetUsersV2)
	r.GET("/api/v2/users/:id", userHandlerV2.GetUsersByIdV2)
	r.POST("/api/v2/users", userHandlerV2.PostUsersV2)
	r.PUT("/api/v2/users/:id", userHandlerV2.PutUsersV2)
	r.DELETE("/api/v2/users/:id", userHandlerV2.DeleteUsersV2)

	r.Run(":8080")
}
