package service

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizService interface {
	Create(ctx context.Context, req dto.QuizDTO) (*dto.QuizDTO, error)
	Update(ctx context.Context, req dto.UpdateQuizDTO) (*dto.QuizDTO, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*dto.QuizDTO, error)
	GetAll(ctx context.Context) ([]*dto.QuizDTO, error)
	GetByCategory(ctx context.Context, category string) ([]*dto.QuizDTO, error)
	GetCategories(ctx context.Context) ([]string, error)
	CountByCategory(ctx context.Context) (map[string]int64, error)
	EnsureIndexes(ctx context.Context) error
}
