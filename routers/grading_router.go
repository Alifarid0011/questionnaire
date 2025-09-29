package routers

import (
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"github.com/gin-gonic/gin"
)

// RegisterGradingRoutes registers routes for grading user answers
func RegisterGradingRoutes(r *gin.Engine, app *wire.App) {

	ratingsRouter := r.Group("/grading/user-answer")
	{
		// Automatically grade a user answer
		ratingsRouter.POST("/:id", app.GradingCtrl.GradeUserAnswer)

		// Manually override score for a specific question
		ratingsRouter.POST("/:id/manual", app.GradingCtrl.ManualGrading)

		// Set appeal flag for a user answer
		ratingsRouter.POST("/:id/appeal", app.GradingCtrl.SetAppeal)
	}
}
