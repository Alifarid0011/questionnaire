package service

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionRatingService interface {
	CreateRating(ctx context.Context, input *dto.QuestionRatingDTO) (*dto.QuestionRatingDTO, error)
	UpdateRating(ctx context.Context, input *dto.UpdateQuestionRatingDTO) (*dto.QuestionRatingDTO, error)
	GetRatingByID(ctx context.Context, id primitive.ObjectID) (*dto.QuestionRatingDTO, error)
	GetRatingsByQuestionID(ctx context.Context, questionID primitive.ObjectID) ([]*dto.QuestionRatingDTO, error)
	GetRatingsByUserID(ctx context.Context, userID primitive.ObjectID) ([]*dto.QuestionRatingDTO, error)
	GetRatingByQuestionAndUser(ctx context.Context, questionID, userID primitive.ObjectID) (*dto.QuestionRatingDTO, error)
}
