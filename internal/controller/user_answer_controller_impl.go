package controller

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto/response"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAnswerController struct {
	service service.UserAnswerService
}

func NewUserAnswerController(s service.UserAnswerService) *UserAnswerController {
	return &UserAnswerController{service: s}
}

func (uac *UserAnswerController) CreateUserAnswer(c *gin.Context) {
	var input dto.UserAnswerDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	answer, err := uac.service.CreateUserAnswer(c.Request.Context(), &input)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(answer).Dispatch()
}

func (uac *UserAnswerController) GetUserAnswerByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	answer, err := uac.service.GetUserAnswerByID(c.Request.Context(), id)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(answer).Dispatch()
}

func (uac *UserAnswerController) GetUserAnswersByQuizID(c *gin.Context) {
	quizIDParam := c.Param("quiz_id")
	quizID, err := primitive.ObjectIDFromHex(quizIDParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	answers, err := uac.service.GetUserAnswersByQuizID(c.Request.Context(), quizID)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(answers).Dispatch()
}

func (uac *UserAnswerController) GetUserAnswersByUserID(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	answers, err := uac.service.GetUserAnswersByUserID(c.Request.Context(), userID)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(answers).Dispatch()
}

func (uac *UserAnswerController) GetUserAnswersByQuizAndUser(c *gin.Context) {
	quizIDParam := c.Param("quiz_id")
	userIDParam := c.Param("user_id")
	quizID, err := primitive.ObjectIDFromHex(quizIDParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	answers, err := uac.service.GetUserAnswersByQuizAndUser(c.Request.Context(), quizID, userID)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(answers).Dispatch()
}
