package repository

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionRatingRepository interface {
	Create(ctx context.Context, rating *models.QuestionRating) error
	Update(ctx context.Context, rating *models.QuestionRating) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.QuestionRating, error)
	FindByQuestionID(ctx context.Context, questionID primitive.ObjectID) ([]*models.QuestionRating, error)
	FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]*models.QuestionRating, error)
	FindByQuestionAndUser(ctx context.Context, questionID, userID primitive.ObjectID) (*models.QuestionRating, error)
	EnsureIndexes(ctx context.Context) error
}
