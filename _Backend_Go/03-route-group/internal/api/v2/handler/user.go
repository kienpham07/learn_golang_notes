package v2handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUsersV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Message": "List all users (V2)"})
}

func (u *UserHandler) GetUsersByIdV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Message": "Get user by id (V2)"})
}

func (u *UserHandler) PostUsersV2(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"Message": "Create user (V2)"})
}

func (u *UserHandler) PutUsersV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Message": "Update user (V2)"})
}

func (u *UserHandler) DeleteUsersV2(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"Message": "Delete user (V2)"})
}
