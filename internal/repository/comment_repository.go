package repository

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *models.Comment) error
	Update(ctx context.Context, comment *models.Comment) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Comment, error)
	FindByQuestionID(ctx context.Context, questionID primitive.ObjectID) ([]*models.Comment, error)
	FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]*models.Comment, error)
	FindReplies(ctx context.Context, parentID primitive.ObjectID) ([]*models.Comment, error)
	EnsureIndexes(ctx context.Context) error
}
