package service

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentService interface {
	CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	UpdateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	GetCommentByID(ctx context.Context, id primitive.ObjectID) (*models.Comment, error)
	GetCommentsByTarget(ctx context.Context, target models.DBRef) ([]*models.Comment, error)
	GetReplies(ctx context.Context, parentID primitive.ObjectID) ([]*models.Comment, error)
	GetCommentsByUser(ctx context.Context, userID primitive.ObjectID) ([]*models.Comment, error)
}
