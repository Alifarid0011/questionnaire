package controller

import "github.com/gin-gonic/gin"

type UserAnswerGradingController interface {
	GradeUserAnswer(c *gin.Context)
	ManualGrading(c *gin.Context)
	SetAppeal(c *gin.Context)
}
