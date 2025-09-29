package service

import (
	"context"
	"errors"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type questionRatingServiceImpl struct {
	repo repository.QuestionRatingRepository
}

func NewQuestionRatingService(repo repository.QuestionRatingRepository) QuestionRatingService {
	return &questionRatingServiceImpl{repo: repo}
}

func (s *questionRatingServiceImpl) CreateRating(ctx context.Context, input *dto.QuestionRatingDTO) (*dto.QuestionRatingDTO, error) {
	existing, _ := s.repo.FindByQuestionAndUser(ctx, input.QuestionID, input.UserID)
	if existing != nil {
		return nil, errors.New("user has already rated this question")
	}

	rating := &models.QuestionRating{
		QuestionID: input.QuestionID,
		UserID:     input.UserID,
		Score:      input.Score,
		CreatedAt:  time.Now(),
	}

	if err := s.repo.Create(ctx, rating); err != nil {
		return nil, err
	}

	return &dto.QuestionRatingDTO{
		ID:         rating.ID,
		QuestionID: rating.QuestionID,
		UserID:     rating.UserID,
		Score:      rating.Score,
		CreatedAt:  rating.CreatedAt,
	}, nil
}

func (s *questionRatingServiceImpl) UpdateRating(ctx context.Context, input *dto.UpdateQuestionRatingDTO) (*dto.QuestionRatingDTO, error) {
	rating, err := s.repo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if input.Score != nil {
		rating.Score = *input.Score
	}
	rating.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, rating); err != nil {
		return nil, err
	}

	return &dto.QuestionRatingDTO{
		ID:         rating.ID,
		QuestionID: rating.QuestionID,
		UserID:     rating.UserID,
		Score:      rating.Score,
		CreatedAt:  rating.CreatedAt,
		UpdatedAt:  rating.UpdatedAt,
	}, nil
}

func (s *questionRatingServiceImpl) GetRatingByID(ctx context.Context, id primitive.ObjectID) (*dto.QuestionRatingDTO, error) {
	rating, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.QuestionRatingDTO{
		ID:         rating.ID,
		QuestionID: rating.QuestionID,
		UserID:     rating.UserID,
		Score:      rating.Score,
		CreatedAt:  rating.CreatedAt,
		UpdatedAt:  rating.UpdatedAt,
	}, nil
}

func (s *questionRatingServiceImpl) GetRatingsByQuestionID(ctx context.Context, questionID primitive.ObjectID) ([]*dto.QuestionRatingDTO, error) {
	ratings, err := s.repo.FindByQuestionID(ctx, questionID)
	if err != nil {
		return nil, err
	}

	var result []*dto.QuestionRatingDTO
	for _, r := range ratings {
		result = append(result, &dto.QuestionRatingDTO{
			ID:         r.ID,
			QuestionID: r.QuestionID,
			UserID:     r.UserID,
			Score:      r.Score,
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
		})
	}
	return result, nil
}

func (s *questionRatingServiceImpl) GetRatingsByUserID(ctx context.Context, userID primitive.ObjectID) ([]*dto.QuestionRatingDTO, error) {
	ratings, err := s.repo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var result []*dto.QuestionRatingDTO
	for _, r := range ratings {
		result = append(result, &dto.QuestionRatingDTO{
			ID:         r.ID,
			QuestionID: r.QuestionID,
			UserID:     r.UserID,
			Score:      r.Score,
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
		})
	}
	return result, nil
}

func (s *questionRatingServiceImpl) GetRatingByQuestionAndUser(ctx context.Context, questionID, userID primitive.ObjectID) (*dto.QuestionRatingDTO, error) {
	rating, err := s.repo.FindByQuestionAndUser(ctx, questionID, userID)
	if err != nil {
		return nil, err
	}
	if rating == nil {
		return nil, nil
	}
	return &dto.QuestionRatingDTO{
		ID:         rating.ID,
		QuestionID: rating.QuestionID,
		UserID:     rating.UserID,
		Score:      rating.Score,
		CreatedAt:  rating.CreatedAt,
		UpdatedAt:  rating.UpdatedAt,
	}, nil
}
