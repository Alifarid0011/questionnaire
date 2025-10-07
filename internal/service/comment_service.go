package service

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentService interface {
	CreateComment(ctx context.Context, comment *dto.CommentDto) (*dto.CommentDto, error)
	UpdateComment(ctx context.Context, comment *dto.CommentDto) (*dto.CommentDto, error)
	GetCommentByID(ctx context.Context, id primitive.ObjectID) (*dto.CommentDto, error)
	GetCommentsByTarget(ctx context.Context, target models.DBRef) ([]*dto.CommentDto, error)
	GetReplies(ctx context.Context, parentID primitive.ObjectID) ([]*dto.CommentDto, error)
	GetCommentsByUser(ctx context.Context, userID primitive.ObjectID) ([]*dto.CommentDto, error)
}
