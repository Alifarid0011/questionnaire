package service

import (
	"context"
	"errors"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserServiceImpl struct {
	userRepo   repository.UserRepository
	casbinRepo repository.CasbinRepository
}

func NewUserService(userRepo repository.UserRepository, casbinRepo repository.CasbinRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: userRepo, casbinRepo: casbinRepo}
}

func (s *UserServiceImpl) FindByUsername(username string, ctx context.Context) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByUsername(username, ctx)
	if err != nil {
		return nil, err
	}
	return mapToUserResponse(user), nil
}

func (s *UserServiceImpl) UpdateUser(uid primitive.ObjectID, req dto.UpdateUserRequest, ctx context.Context) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByUID(uid, ctx)
	if err != nil {
		return nil, err
	}
	if errUpdateStruct := utils.UpdateStruct(user, req); errUpdateStruct != nil {
		return nil, errUpdateStruct
	}
	if errUpdate := s.userRepo.Update(user, ctx); errUpdate != nil {
		return nil, errUpdate
	}
	return mapToUserResponse(user), nil
}

func (s *UserServiceImpl) FindByUID(uid primitive.ObjectID, ctx context.Context) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByUID(uid, ctx)
	if err != nil {
		return nil, err
	}
	return mapToUserResponse(user), nil
}

func (s *UserServiceImpl) CreateUser(req dto.CreateUserRequest, ctx context.Context) (*dto.UserResponse, error) {
	_, err := s.userRepo.FindByUsername(req.Username, ctx)
	if err == nil {
		return nil, errors.New("username already exists")
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		UID:          utils.GenerateUID(),
		Username:     req.Username,
		Email:        req.Email,
		NationalCode: req.NationalCode,
		Password:     hashedPassword,
		Mobile:       req.Mobile,
		FullName:     req.FullName,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if errCreate := s.userRepo.Create(user, ctx); errCreate != nil {
		return nil, errCreate
	}
	_, errAddGroupingPolicy := s.casbinRepo.AddGroupingPolicy(user.UID.Hex(), constant.RoleUser)
	if errAddGroupingPolicy != nil {
		return nil, errAddGroupingPolicy
	}
	return mapToUserResponse(user), nil
}

func (s *UserServiceImpl) GetAll(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := s.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var result []dto.UserResponse
	for _, u := range users {
		result = append(result, *mapToUserResponse(&u))
	}
	return result, nil
}

func (s *UserServiceImpl) DeleteUser(uid primitive.ObjectID, ctx context.Context) error {
	return errors.New("delete not implemented in repository yet")
}

func (s *UserServiceImpl) Me(userID primitive.ObjectID, ctx context.Context) (*dto.UserResponse, error) {
	return s.FindByUID(userID, ctx)
}

func mapToUserResponse(u *models.User) *dto.UserResponse {
	return &dto.UserResponse{
		UID:          u.UID,
		Username:     u.Username,
		Email:        u.Email,
		FullName:     u.FullName,
		Mobile:       u.Mobile,
		NationalCode: u.NationalCode,
		CreatedAt:    u.CreatedAt,
	}
}
