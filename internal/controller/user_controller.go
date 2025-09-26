package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	FindByUsername(ctx *gin.Context)
	FindUserByUID(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	Me(ctx *gin.Context)
}
