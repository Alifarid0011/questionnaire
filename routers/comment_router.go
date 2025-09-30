package routers

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/middleware"
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(r *gin.Engine, app *wire.App) {
	commentRouter := r.Group("/comments",
		middleware.AuthMiddleware(app.BlackListRepo, app.TokenManager),
		middleware.CasbinMiddleware(app.Enforcer))
	{
		commentRouter.POST("", app.CommentCtrl.CreateComment)                            // Create a new comment
		commentRouter.PUT("", app.CommentCtrl.UpdateComment)                             // Update a comment
		commentRouter.GET("/:id", app.CommentCtrl.GetCommentByID)                        // Get comment by ID
		commentRouter.GET("/question/:question_id", app.CommentCtrl.GetCommentsByTarget) // Get all comments for a target (question)
		commentRouter.GET("/user/:user_id", app.CommentCtrl.GetCommentsByUser)           // Get all comments by a user
		commentRouter.GET("/parent/:parent_id", app.CommentCtrl.GetReplies)              // Get replies for a parent comment
	}
}
