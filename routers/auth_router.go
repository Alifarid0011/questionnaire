package routers

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/middleware"
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, app *wire.App) {
	auth := r.Group("/auth", middleware.UserAgentMiddleware(), middleware.UserAgentMiddleware())
	{
		auth.POST("/login", app.AuthCtrl.Login)
		auth.POST("/refresh_token", app.AuthCtrl.UseRefreshToken)
		auth.GET("/logout", middleware.AuthMiddleware(app.BlackListRepo, app.TokenManager), app.AuthCtrl.Logout)
		auth.POST("/register", app.AuthCtrl.Register)
	}
}
