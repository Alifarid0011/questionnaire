package controller

import "github.com/gin-gonic/gin"

type QuizControllerInterface interface {
	CreateQuiz(c *gin.Context)
	UpdateQuiz(c *gin.Context)
	DeleteQuiz(c *gin.Context)
	GetQuizByID(c *gin.Context)
	GetAllQuizzes(c *gin.Context)
	GetQuizzesByCategory(c *gin.Context)
	GetQuizCategories(c *gin.Context)
	GetQuizCountByCategory(c *gin.Context)
}
