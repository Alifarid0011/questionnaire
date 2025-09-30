package routers

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/middleware"
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, app *wire.App) {
	userRouter := r.Group("/users",
		middleware.AuthMiddleware(app.BlackListRepo, app.TokenManager),
		middleware.CasbinMiddleware(app.Enforcer))
	{
		userRouter.POST("", app.UserCtrl.CreateUser)
		userRouter.GET("", middleware.PaginationMiddleware(), app.UserCtrl.GetAllUsers)
		userRouter.GET("/me", app.UserCtrl.Me)
		userRouter.GET("/uid/:uid", app.UserCtrl.FindUserByUID)
		userRouter.GET("/username/:username", app.UserCtrl.FindByUsername)
		userRouter.PUT("/:uid", app.UserCtrl.UpdateUser)
		userRouter.DELETE("/:uid", app.UserCtrl.DeleteUser)
	}
}
