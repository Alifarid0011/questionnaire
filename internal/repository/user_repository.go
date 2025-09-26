package repository

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	FindByUsername(username string, ctx context.Context) (*models.User, error)
	FindByUID(uid primitive.ObjectID, ctx context.Context) (*models.User, error)
	Create(user *models.User, ctx context.Context) error
	GetAll(ctx context.Context) ([]models.User, error)
	Update(user *models.User, ctx context.Context) error // Add Update
	Delete(user *models.User, ctx context.Context) error // Add Delete
	EnsureIndexes() error
}
