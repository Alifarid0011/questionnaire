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

// CreateRating godoc
// @Summary Create a new question rating
// @Description Allows a user to rate a question from 1 to 5
// @Tags QuestionRatings
// @Accept json
// @Produce json
// @Security AuthBearer
// @Param rating body dto.QuestionRatingDTO true "Question Rating"
// @Success 200 {object} response.Response{data=dto.QuestionRatingDTO}
// @Failure 400 {object} response.Response
// @Router /ratings [post]
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

// UpdateRating godoc
// @Summary Update an existing question rating
// @Description Update the score of a previously submitted rating
// @Tags QuestionRatings
// @Accept json
// @Produce json
// @Security AuthBearer
// @Param rating body dto.UpdateQuestionRatingDTO true "Update Question Rating"
// @Success 200 {object} response.Response{data=dto.QuestionRatingDTO}
// @Failure 400 {object} response.Response
// @Router /ratings [put]
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

// GetRatingByID godoc
// @Summary Get a question rating by ID
// @Description Retrieve a specific rating by its ID
// @Tags QuestionRatings
// @Produce json
// @Security AuthBearer
// @Param id path string true "Rating ID"
// @Success 200 {object} response.Response{data=dto.QuestionRatingDTO}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /ratings/{id} [get]
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

// GetRatingsByQuestionID godoc
// @Summary Get all ratings for a question
// @Description Retrieve all ratings submitted for a specific question
// @Tags QuestionRatings
// @Security AuthBearer
// @Produce json
// @Param question_id path string true "Question ID"
// @Success 200 {array} dto.QuestionRatingDTO
// @Failure 400 {object} response.Response
// @Router /ratings/question/{question_id} [get]
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

// GetRatingsByUserID godoc
// @Summary Get all ratings by a user
// @Description Retrieve all question ratings submitted by a specific user
// @Tags QuestionRatings
// @Security AuthBearer
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {array} dto.QuestionRatingDTO
// @Failure 400 {object} response.Response
// @Router /ratings/user/{user_id} [get]
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

// GetRatingByQuestionAndUser godoc
// @Summary Get a user's rating for a specific question
// @Description Retrieve the rating submitted by a specific user for a specific question
// @Tags QuestionRatings
// @Security AuthBearer
// @Produce json
// @Param question_id path string true "Question ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} dto.QuestionRatingDTO
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /ratings/question/{question_id}/user/{user_id} [get]
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
