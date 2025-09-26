package controller

import "github.com/gin-gonic/gin"

type CasbinController interface {
	CheckPermission(ctx *gin.Context)
	CreatePolicy(ctx *gin.Context)
	RemovePolicy(ctx *gin.Context)
	AddGroupingPolicy(ctx *gin.Context)
	RemoveGroupingPolicy(ctx *gin.Context)
	ListAllCasbinData(ctx *gin.Context)
	PermissionsTree(ctx *gin.Context)
	Roles(ctx *gin.Context)
	UserRoles(ctx *gin.Context)
}
