package controller

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto/response"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizControllerImpl struct {
	service service.QuizService
}

func NewQuizController(s service.QuizService) QuizController {
	return &QuizControllerImpl{service: s}
}

// CreateQuiz godoc
// @Summary Create a new quiz
// @Description Creates a new quiz with questions
// @Tags quizzes
// @Accept json
// @Produce json
// @Security AuthBearer
// @Param quiz body dto.QuizDTO true "Quiz data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /quizzes [post]
func (qc *QuizControllerImpl) CreateQuiz(c *gin.Context) {
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

// UpdateQuiz godoc
// @Summary Update an existing quiz
// @Description Updates quiz details by ID
// @Tags quizzes
// @Accept json
// @Produce json
// @Security AuthBearer
// @Param quiz body dto.UpdateQuizDTO true "Updated quiz data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /quizzes [put]
func (qc *QuizControllerImpl) UpdateQuiz(c *gin.Context) {
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

// DeleteQuiz godoc
// @Summary Delete a quiz
// @Description Deletes a quiz by ID
// @Tags quizzes
// @Security AuthBearer
// @Param id path string true "Quiz ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /quizzes/{id} [delete]
func (qc *QuizControllerImpl) DeleteQuiz(c *gin.Context) {
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

// GetQuizByID godoc
// @Summary Get quiz by ID
// @Description Returns a single quiz by ID
// @Tags quizzes
// @Security AuthBearer
// @Param id path string true "Quiz ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /quizzes/{id} [get]
func (qc *QuizControllerImpl) GetQuizByID(c *gin.Context) {
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

// GetAllQuizzes godoc
// @Summary Get all quizzes
// @Description Returns all quizzes
// @Tags quizzes
// @Security AuthBearer
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /quizzes [get]
func (qc *QuizControllerImpl) GetAllQuizzes(c *gin.Context) {
	quizzes, err := qc.service.GetAll(c.Request.Context())
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(quizzes).Dispatch()
}

// GetQuizzesByCategory godoc
// @Summary Get quizzes by category
// @Description Returns all quizzes for a given category
// @Tags quizzes
// @Security AuthBearer
// @Param category query string true "Category name"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /quizzes/category [get]
func (qc *QuizControllerImpl) GetQuizzesByCategory(c *gin.Context) {
	category := c.Query("category")
	quizzes, err := qc.service.GetByCategory(c.Request.Context(), category)
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(quizzes).Dispatch()
}

// GetQuizCategories godoc
// @Summary Get quiz categories
// @Description Returns all available quiz categories
// @Tags quizzes
// @Security AuthBearer
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /quizzes/categories [get]
func (qc *QuizControllerImpl) GetQuizCategories(c *gin.Context) {
	categories, err := qc.service.GetCategories(c.Request.Context())
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(categories).Dispatch()
}

// GetQuizCountByCategory godoc
// @Summary Get quiz count by category
// @Description Returns number of quizzes in each category
// @Tags quizzes
// @Security AuthBearer
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /quizzes/categories/count [get]
func (qc *QuizControllerImpl) GetQuizCountByCategory(c *gin.Context) {
	countMap, err := qc.service.CountByCategory(c.Request.Context())
	if err != nil {
		response.New(c).Errors(err).Dispatch()
		return
	}
	response.New(c).Data(countMap).Dispatch()
}
