package service

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type userAnswerServiceImpl struct {
	repo repository.UserAnswerRepository
}

func NewUserAnswerService(repo repository.UserAnswerRepository) UserAnswerService {
	return &userAnswerServiceImpl{repo: repo}
}

func (s *userAnswerServiceImpl) CreateUserAnswer(ctx context.Context, input *dto.UserAnswerDTO) (*dto.UserAnswerDTO, error) {
	answer := &models.UserAnswer{
		ID:        primitive.NewObjectID(),
		QuizID:    input.QuizID,
		UserID:    input.UserID,
		Answers:   input.Answers,
		CreatedAt: time.Now(),
	}
	if err := s.repo.UserAnswerCreate(ctx, answer); err != nil {
		return nil, err
	}
	return &dto.UserAnswerDTO{
		QuizID:  answer.QuizID,
		UserID:  answer.UserID,
		Answers: answer.Answers,
	}, nil
}

func (s *userAnswerServiceImpl) GetUserAnswerByID(ctx context.Context, id primitive.ObjectID) (*dto.UserAnswerDTO, error) {
	answer, err := s.repo.UserAnswerFindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.UserAnswerDTO{
		QuizID:  answer.QuizID,
		UserID:  answer.UserID,
		Answers: answer.Answers,
	}, nil
}

func (s *userAnswerServiceImpl) GetUserAnswersByQuizID(ctx context.Context, quizID primitive.ObjectID) ([]*dto.UserAnswerDTO, error) {
	answers, err := s.repo.UserAnswerFindByQuizID(ctx, quizID)
	if err != nil {
		return nil, err
	}
	result := make([]*dto.UserAnswerDTO, len(answers))
	for i, ans := range answers {
		result[i] = &dto.UserAnswerDTO{
			QuizID:  ans.QuizID,
			UserID:  ans.UserID,
			Answers: ans.Answers,
		}
	}
	return result, nil
}

func (s *userAnswerServiceImpl) GetUserAnswersByUserID(ctx context.Context, userID primitive.ObjectID) ([]*dto.UserAnswerDTO, error) {
	answers, err := s.repo.UserAnswerFindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	result := make([]*dto.UserAnswerDTO, len(answers))
	for i, ans := range answers {
		result[i] = &dto.UserAnswerDTO{
			QuizID:  ans.QuizID,
			UserID:  ans.UserID,
			Answers: ans.Answers,
		}
	}
	return result, nil
}

func (s *userAnswerServiceImpl) GetUserAnswersByQuizAndUser(ctx context.Context, quizID, userID primitive.ObjectID) ([]*dto.UserAnswerDTO, error) {
	answers, err := s.repo.UserAnswerFindByQuizIDAndUserID(ctx, quizID, userID)
	if err != nil {
		return nil, err
	}
	result := make([]*dto.UserAnswerDTO, len(answers))
	for i, ans := range answers {
		result[i] = &dto.UserAnswerDTO{
			QuizID:  ans.QuizID,
			UserID:  ans.UserID,
			Answers: ans.Answers,
		}
	}
	return result, nil
}
