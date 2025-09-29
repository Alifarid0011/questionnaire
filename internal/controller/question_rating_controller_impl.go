package controller

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto/response"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type questionRatingControllerImpl struct {
	service service.QuestionRatingService
}

func NewQuestionRatingController(s service.QuestionRatingService) QuestionRatingController {
	return &questionRatingControllerImpl{service: s}
}

func (qc *questionRatingControllerImpl) CreateRating(c *gin.Context) {
	var input dto.QuestionRatingDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	rating, err := qc.service.CreateRating(c.Request.Context(), &input)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(rating).Dispatch()
}

func (qc *questionRatingControllerImpl) UpdateRating(c *gin.Context) {
	var input dto.UpdateQuestionRatingDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	rating, err := qc.service.UpdateRating(c.Request.Context(), &input)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(rating).Dispatch()
}

func (qc *questionRatingControllerImpl) GetRatingByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	rating, err := qc.service.GetRatingByID(c.Request.Context(), id)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(rating).Dispatch()
}

func (qc *questionRatingControllerImpl) GetRatingsByQuestionID(c *gin.Context) {
	questionIDParam := c.Param("question_id")
	questionID, err := primitive.ObjectIDFromHex(questionIDParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	ratings, err := qc.service.GetRatingsByQuestionID(c.Request.Context(), questionID)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(ratings).Dispatch()
}

func (qc *questionRatingControllerImpl) GetRatingsByUserID(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	ratings, err := qc.service.GetRatingsByUserID(c.Request.Context(), userID)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(ratings).Dispatch()
}

func (qc *questionRatingControllerImpl) GetRatingByQuestionAndUser(c *gin.Context) {
	questionIDParam := c.Param("question_id")
	userIDParam := c.Param("user_id")

	questionID, err := primitive.ObjectIDFromHex(questionIDParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	rating, err := qc.service.GetRatingByQuestionAndUser(c.Request.Context(), questionID, userID)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(rating).Dispatch()
}
