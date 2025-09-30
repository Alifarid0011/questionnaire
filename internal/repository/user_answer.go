package repository

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAnswerRepository interface {
	UserAnswerCreate(ctx context.Context, answer *models.UserAnswer) error
	UserAnswerFindByID(ctx context.Context, id primitive.ObjectID) (*models.UserAnswer, error)
	UserAnswerFindByQuizID(ctx context.Context, quizID primitive.ObjectID) ([]*models.UserAnswer, error)
	UserAnswerFindByUserID(ctx context.Context, userID primitive.ObjectID) ([]*models.UserAnswer, error)
	UserAnswerFindByQuizIDAndUserID(ctx context.Context, quizID, userID primitive.ObjectID) ([]*models.UserAnswer, error)
	EnsureIndexes(ctx context.Context) error
	UserAnswerUpdate(ctx context.Context, answer *models.UserAnswer) error
	UserAnswerSetAppeal(ctx context.Context, id primitive.ObjectID, appeal bool) error
}
