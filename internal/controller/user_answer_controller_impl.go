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

// CreateUserAnswer godoc
// @Summary Create a new user answer
// @Description Create and store a user answer for a quiz
// @Tags UserAnswers
// @Accept json
// @Produce json
// @Param userAnswer body dto.UserAnswerDTO true "User Answer DTO"
// @Success 200 {object} response.Response{data=dto.UserAnswerDTO}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /user-answers [post]
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

// GetUserAnswerByID godoc
// @Summary Get user answer by ID
// @Description Retrieve a user answer using its ID
// @Tags UserAnswers
// @Produce json
// @Param id path string true "User Answer ID"
// @Success 200 {object} response.Response{data=dto.UserAnswerDTO}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /user-answers/{id} [get]
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

// GetUserAnswersByQuizID godoc
// @Summary Get all user answers for a quiz
// @Description Retrieve all user answers for a specific quiz
// @Tags UserAnswers
// @Produce json
// @Param quiz_id path string true "Quiz ID"
// @Success 200 {object} response.Response{data=[]dto.UserAnswerDTO}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /user-answers/quiz/{quiz_id} [get]
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

// GetUserAnswersByUserID godoc
// @Summary Get all user answers for a user
// @Description Retrieve all user answers submitted by a specific user
// @Tags UserAnswers
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} response.Response{data=[]dto.UserAnswerDTO}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /user-answers/user/{user_id} [get]
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

// GetUserAnswersByQuizAndUser godoc
// @Summary Get user answers for a quiz and user
// @Description Retrieve answers submitted by a specific user for a specific quiz
// @Tags UserAnswers
// @Produce json
// @Param quiz_id path string true "Quiz ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} response.Response{data=[]dto.UserAnswerDTO}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /user-answers/quiz/{quiz_id}/user/{user_id} [get]
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
