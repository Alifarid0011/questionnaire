package routers

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/middleware"
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"github.com/gin-gonic/gin"
)

func RegisterQuestionRatingRoutes(r *gin.Engine, app *wire.App) {
	ratingsRouter := r.Group("/ratings",
		middleware.AuthMiddleware(app.BlackListRepo, app.TokenManager),
		middleware.CasbinMiddleware(app.Enforcer),
	)
	{
		ratingsRouter.POST("", app.QuestionRatingCtrl.CreateRating)                                                  // Create a new rating
		ratingsRouter.PUT("", app.QuestionRatingCtrl.UpdateRating)                                                   // Update a rating
		ratingsRouter.GET("/:id", app.QuestionRatingCtrl.GetRatingByID)                                              // Get rating by ID
		ratingsRouter.GET("/question/:question_id", app.QuestionRatingCtrl.GetRatingsByQuestionID)                   // Get all ratings for a question
		ratingsRouter.GET("/user/:user_id", app.QuestionRatingCtrl.GetRatingsByUserID)                               // Get all ratings by a user
		ratingsRouter.GET("/question/:question_id/user/:user_id", app.QuestionRatingCtrl.GetRatingByQuestionAndUser) // Get rating by question and user
	}
}
