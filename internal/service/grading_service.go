package service

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GradingService interface {
	GradeUserAnswer(ctx context.Context, ua *models.UserAnswer) error
	ManualGrading(ctx context.Context, uaID primitive.ObjectID, questionID string, newScore float64) error
	GradeUserAnswerByID(ctx context.Context, uaID primitive.ObjectID) (*models.UserAnswer, error)
	ManualGradingByID(ctx context.Context, uaID primitive.ObjectID, questionID string, newScore float64) (*models.UserAnswer, error)
	SetAppeal(ctx context.Context, uaID primitive.ObjectID, appeal bool) error
}
