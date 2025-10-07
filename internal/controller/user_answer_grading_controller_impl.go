package controller

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto/response"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAnswerGradingControllerImpl struct {
	gradingService service.GradingService
}

func NewUserAnswerGradingController(gs service.GradingService) *UserAnswerGradingControllerImpl {
	return &UserAnswerGradingControllerImpl{
		gradingService: gs,
	}
}

// GradeUserAnswer godoc
// @Summary Automatically grade a user answer
// @Description Grade all questions in a user answer using automatic grading
// @Tags Grading
// @Security AuthBearer
// @Accept json
// @Produce json
// @Param id path string true "User Answer ID"
// @Success 200 {object} response.Response{data=models.UserAnswer}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /grading/user-answer/{id} [post]
func (uc *UserAnswerGradingControllerImpl) GradeUserAnswer(c *gin.Context) {
	idParam := c.Param("id")
	uaID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	ua, err := uc.gradingService.GradeUserAnswerByID(c.Request.Context(), uaID)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	response.New(c).Data(ua).Dispatch()
}

// ManualGrading godoc
// @Summary Override score for a specific question
// @Description Allows a grader to manually change the score of a single question
// @Tags Grading
// @Security AuthBearer
// @Accept json
// @Produce json
// @Param id path string true "User Answer ID"
// @Success 200 {object} response.Response{data=models.UserAnswer}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /grading/user-answer/{id}/manual [post]
func (uc *UserAnswerGradingControllerImpl) ManualGrading(c *gin.Context) {
	idParam := c.Param("id")
	uaID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	var body struct {
		QuestionID primitive.ObjectID `json:"question_id"`
		NewScore   float64            `json:"new_score"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	ua, err := uc.gradingService.ManualGradingByID(c.Request.Context(), uaID, body.QuestionID, body.NewScore)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	response.New(c).Data(ua).Dispatch()
}

// SetAppeal godoc
// @Summary Set appeal flag for a user answer
// @Description Allows a user to mark their answer as appealed for manual review
// @Tags Grading
// @Security AuthBearer
// @Accept json
// @Produce json
// @Param id path string true "User Answer ID"
// @Success 200 {object} response.Response{data=map[string]bool}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /grading/user-answer/{id}/appeal [post]
func (uc *UserAnswerGradingControllerImpl) SetAppeal(c *gin.Context) {
	idParam := c.Param("id")
	uaID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	var body struct {
		Appeal bool `json:"appeal"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	if err := uc.gradingService.SetAppeal(c.Request.Context(), uaID, body.Appeal); err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}

	response.New(c).Data(map[string]bool{"appeal": body.Appeal}).Dispatch()
}
