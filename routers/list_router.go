package routers

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/middleware"
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"github.com/gin-gonic/gin"
)

func RegisterListRoutes(r *gin.Engine, app *wire.App) {
	router := r.Group("/routes", middleware.UserAgentMiddleware())
	{
		router.GET("/list", app.RouterCtr.ListGroupedRoutes)
	}
}
