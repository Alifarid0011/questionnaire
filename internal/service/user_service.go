package service

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	FindByUsername(username string, ctx context.Context) (*dto.UserResponse, error)
	FindByUID(uid primitive.ObjectID, ctx context.Context) (*dto.UserResponse, error)
	CreateUser(req dto.CreateUserRequest, ctx context.Context) (*dto.UserResponse, error)
	GetAll(ctx context.Context) ([]dto.UserResponse, error)
	UpdateUser(uid primitive.ObjectID, req dto.UpdateUserRequest, ctx context.Context) (*dto.UserResponse, error)
	DeleteUser(uid primitive.ObjectID, ctx context.Context) error
	Me(userID primitive.ObjectID, ctx context.Context) (*dto.UserResponse, error)
}
