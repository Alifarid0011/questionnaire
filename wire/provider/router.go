package provider

import "github.com/gin-gonic/gin"

func RouterEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	return r
}
