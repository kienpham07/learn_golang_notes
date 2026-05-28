package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/demo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello world", "number": 67})
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "List of users"})
	})

	r.GET("/user/:user_id", func(ctx *gin.Context) {
		user_id := ctx.Param("user_id") // Return string
		ctx.JSON(200, gin.H{
			"data":    "User information",
			"user_id": user_id,
		})
	})

	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "List of products"})
	})

	r.GET("/product/:product_name", func(ctx *gin.Context) {
		product_name := ctx.Param("product_name")
		price := ctx.Query("price")
		color := ctx.Query("color") // Query parameter is optional

		ctx.JSON(200, gin.H{
			"data":         "Product information",
			"product_name": product_name,
			"price":        price,
			"color":        color,
		})
	})

	r.Run(":8080")
}
