package routers

import (
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"github.com/gin-gonic/gin"
)

func RegisterQuizRoutes(r *gin.Engine, app *wire.App) {
	quizRouter := r.Group("/quizzes")
	{
		quizRouter.POST("", app.QuizCtrl.CreateQuiz)
		quizRouter.PUT("", app.QuizCtrl.UpdateQuiz)
		quizRouter.DELETE("/:id", app.QuizCtrl.DeleteQuiz)
		quizRouter.GET("/:id", app.QuizCtrl.GetQuizByID)
		quizRouter.GET("", app.QuizCtrl.GetAllQuizzes)
		quizRouter.GET("/category", app.QuizCtrl.GetQuizzesByCategory)
		quizRouter.GET("/categories", app.QuizCtrl.GetQuizCategories)
		quizRouter.GET("/categories/count", app.QuizCtrl.GetQuizCountByCategory)
	}
}
