package controller

import "github.com/gin-gonic/gin"

type QuestionRatingController interface {
	CreateRating(c *gin.Context)
	UpdateRating(c *gin.Context)
	GetRatingByID(c *gin.Context)
	GetRatingsByQuestionID(c *gin.Context)
	GetRatingsByUserID(c *gin.Context)
	GetRatingByQuestionAndUser(c *gin.Context)
}
