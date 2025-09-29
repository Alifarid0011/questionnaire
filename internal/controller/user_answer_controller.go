package controller

import "github.com/gin-gonic/gin"

type UserAnswerController interface {
	CreateUserAnswer(c *gin.Context)
	GetUserAnswerByID(c *gin.Context)
	GetUserAnswersByQuizID(c *gin.Context)
	GetUserAnswersByUserID(c *gin.Context)
	GetUserAnswersByQuizAndUser(c *gin.Context)
}
