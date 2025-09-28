package controller

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto/response"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizController struct {
	service service.QuizService
}

func NewQuizController(s service.QuizService) *QuizController {
	return &QuizController{service: s}
}

func (qc *QuizController) CreateQuiz(c *gin.Context) {
	var input dto.QuizDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	quiz, err := qc.service.Create(c.Request.Context(), input)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(quiz).Dispatch()
}

func (qc *QuizController) UpdateQuiz(c *gin.Context) {
	var input dto.UpdateQuizDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	quiz, err := qc.service.Update(c.Request.Context(), input)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(quiz).Dispatch()
}

func (qc *QuizController) DeleteQuiz(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	if err := qc.service.Delete(c.Request.Context(), id); err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Message("Quiz deleted successfully").Dispatch()
}

func (qc *QuizController) GetQuizByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	quiz, err := qc.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(quiz).Dispatch()
}

func (qc *QuizController) GetAllQuizzes(c *gin.Context) {
	quizzes, err := qc.service.GetAll(c.Request.Context())
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(quizzes).Dispatch()
}

func (qc *QuizController) GetQuizzesByCategory(c *gin.Context) {
	category := c.Query("category")
	quizzes, err := qc.service.GetByCategory(c.Request.Context(), category)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(quizzes).Dispatch()
}

func (qc *QuizController) GetQuizCategories(c *gin.Context) {
	categories, err := qc.service.GetCategories(c.Request.Context())
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(categories).Dispatch()
}

func (qc *QuizController) GetQuizCountByCategory(c *gin.Context) {
	countMap, err := qc.service.CountByCategory(c.Request.Context())
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(countMap).Dispatch()
}
