package routers

import (
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"github.com/gin-gonic/gin"
)

func RegisterUserAnswerRoutes(r *gin.Engine, app *wire.App) {
	userAnswerRouter := r.Group("/user-answers")
	{
		userAnswerRouter.POST("", app.UserAnswerCtrl.CreateUserAnswer)
		userAnswerRouter.GET("/:id", app.UserAnswerCtrl.GetUserAnswerByID)
		userAnswerRouter.GET("/quiz/:quiz_id", app.UserAnswerCtrl.GetUserAnswersByQuizID)
		userAnswerRouter.GET("/user/:user_id", app.UserAnswerCtrl.GetUserAnswersByUserID)
		userAnswerRouter.GET("/quiz/:quiz_id/user/:user_id", app.UserAnswerCtrl.GetUserAnswersByQuizAndUser)
	}
}
