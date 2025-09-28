package service

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAnswerService interface {
	CreateUserAnswer(ctx context.Context, input *dto.UserAnswerDTO) (*dto.UserAnswerDTO, error)
	GetUserAnswerByID(ctx context.Context, id primitive.ObjectID) (*dto.UserAnswerDTO, error)
	GetUserAnswersByQuizID(ctx context.Context, quizID primitive.ObjectID) ([]*dto.UserAnswerDTO, error)
	GetUserAnswersByUserID(ctx context.Context, userID primitive.ObjectID) ([]*dto.UserAnswerDTO, error)
	GetUserAnswersByQuizAndUser(ctx context.Context, quizID, userID primitive.ObjectID) ([]*dto.UserAnswerDTO, error)
}
