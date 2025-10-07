package service

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type quizServiceImpl struct {
	repo repository.QuizRepository
}

func NewQuizService(repo repository.QuizRepository) QuizService {
	return &quizServiceImpl{repo: repo}
}

func (s *quizServiceImpl) Create(ctx context.Context, req dto.QuizDTO) (*dto.QuizDTO, error) {
	quiz := &models.Quiz{
		ID:        primitive.NewObjectID(),
		UserID:    ctx.Value("user_uid").(primitive.ObjectID),
		Title:     req.Title,
		Category:  req.Category,
		Level:     req.Level,
		Questions: mapQuestionsDTOToModel(req.Questions),
	}
	if err := s.repo.QuizCreate(ctx, quiz); err != nil {
		return nil, err
	}
	return mapQuizModelToDTO(quiz), nil
}

func (s *quizServiceImpl) Update(ctx context.Context, req dto.UpdateQuizDTO) (*dto.QuizDTO, error) {
	quiz, err := s.repo.QuizFindByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if req.Title != nil {
		quiz.Title = *req.Title
	}
	if req.Category != nil {
		quiz.Category = *req.Category
	}
	if req.Level != nil {
		quiz.Level = *req.Level
	}
	if req.Questions != nil {
		quiz.Questions = mapQuestionsDTOToModel(*req.Questions)
	}
	if err := s.repo.QuizUpdate(ctx, quiz); err != nil {
		return nil, err
	}
	return mapQuizModelToDTO(quiz), nil
}

func (s *quizServiceImpl) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.QuizDelete(ctx, id)
}

func (s *quizServiceImpl) GetByID(ctx context.Context, id primitive.ObjectID) (*dto.QuizDTO, error) {
	quiz, err := s.repo.QuizFindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapQuizModelToDTO(quiz), nil
}

func (s *quizServiceImpl) GetAll(ctx context.Context) ([]*dto.QuizDTO, error) {
	quizzes, err := s.repo.QuizGetAll(ctx)
	if err != nil {
		return nil, err
	}
	return mapQuizModelsToDTOs(quizzes), nil
}

func (s *quizServiceImpl) GetByCategory(ctx context.Context, category string) ([]*dto.QuizDTO, error) {
	quizzes, err := s.repo.QuizGetByCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return mapQuizModelsToDTOs(quizzes), nil
}

func (s *quizServiceImpl) GetCategories(ctx context.Context) ([]string, error) {
	return s.repo.QuizGetCategories(ctx)
}

func (s *quizServiceImpl) CountByCategory(ctx context.Context) (map[string]int64, error) {
	return s.repo.QuizCountByCategory(ctx)
}

func mapQuestionsDTOToModel(dtos []dto.QuestionDTO) []models.Question {
	questions := make([]models.Question, len(dtos))
	for i, q := range dtos {
		questions[i] = models.Question{
			ID:            primitive.NewObjectID(),
			Type:          q.Type,
			Label:         q.Label,
			Options:       q.Options,
			CorrectAnswer: q.CorrectAnswer,
			KeyWords:      q.KeyWords,
		}
	}
	return questions
}

func mapQuizModelToDTO(q *models.Quiz) *dto.QuizDTO {
	questions := make([]dto.QuestionDTO, len(q.Questions))
	for i, qs := range q.Questions {
		questions[i] = dto.QuestionDTO{
			ID:            qs.ID,
			Type:          qs.Type,
			Label:         qs.Label,
			Options:       qs.Options,
			CorrectAnswer: qs.CorrectAnswer,
			KeyWords:      qs.KeyWords,
		}
	}
	return &dto.QuizDTO{
		ID:        q.ID,
		Title:     q.Title,
		Category:  q.Category,
		Level:     q.Level,
		UserID:    &q.UserID,
		Questions: questions,
		CreatedAt: q.CreatedAt,
	}
}

func mapQuizModelsToDTOs(quizzes []*models.Quiz) []*dto.QuizDTO {
	res := make([]*dto.QuizDTO, len(quizzes))
	for i, q := range quizzes {
		res[i] = mapQuizModelToDTO(q)
	}
	return res
}
