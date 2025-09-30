package repository

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizRepository interface {
	QuizCreate(ctx context.Context, quiz *models.Quiz) error
	QuizFindByID(ctx context.Context, id primitive.ObjectID) (*models.Quiz, error)
	QuizUpdate(ctx context.Context, quiz *models.Quiz) error
	QuizDelete(ctx context.Context, id primitive.ObjectID) error
	QuizGetAll(ctx context.Context) ([]*models.Quiz, error)
	QuizGetByCategory(ctx context.Context, category string) ([]*models.Quiz, error)
	QuizGetCategories(ctx context.Context) ([]string, error)
	QuizCountByCategory(ctx context.Context) (map[string]int64, error)
	EnsureIndexes(ctx context.Context) error
}
