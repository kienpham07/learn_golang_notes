package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (u *ProductHandler) GetProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Message": "List all products (V1)"})
}

func (u *ProductHandler) GetProductsByIdV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Message": "Get product by id (V1)"})
}

func (u *ProductHandler) PostProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"Message": "Create product (V1)"})
}

func (u *ProductHandler) PutProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Message": "Update product (V1)"})
}

func (u *ProductHandler) DeleteProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"Message": "Delete product (V1)"})
}
