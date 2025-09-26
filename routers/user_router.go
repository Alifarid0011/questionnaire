package router

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/middleware"
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, app *wire.App) {
	userRouter := r.Group("/users")
	{
		userRouter.POST("", app.UserCtrl.Create)
		userRouter.GET("", middleware.PaginationMiddleware(), app.UserCtrl.GetAll)
		userRouter.GET("/me", app.UserCtrl.Me)
		userRouter.GET("/uid/:uid", app.UserCtrl.FindByUID)
		userRouter.GET("/username/:username", app.UserCtrl.FindByUsername)
		userRouter.PUT("/:uid", app.UserCtrl.Update)
		userRouter.DELETE("/:uid", app.UserCtrl.Delete)
	}
}
