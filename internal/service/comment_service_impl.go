package service

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type commentServiceImpl struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentServiceImpl{repo: repo}
}

// CreateComment creates a new comment
func (s *commentServiceImpl) CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	comment.CreatedAt = time.Now()
	if err := s.repo.Create(ctx, comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// UpdateComment updates an existing comment
func (s *commentServiceImpl) UpdateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	comment.UpdatedAt = time.Now()
	if err := s.repo.Update(ctx, comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// GetCommentByID returns a comment by its ID
func (s *commentServiceImpl) GetCommentByID(ctx context.Context, id primitive.ObjectID) (*models.Comment, error) {
	return s.repo.FindByID(ctx, id)
}

// GetCommentsByTarget returns comments for a polymorphic target
func (s *commentServiceImpl) GetCommentsByTarget(ctx context.Context, target models.DBRef) ([]*models.Comment, error) {
	return s.repo.FindByTarget(ctx, target)
}

// GetReplies returns replies to a parent comment
func (s *commentServiceImpl) GetReplies(ctx context.Context, parentID primitive.ObjectID) ([]*models.Comment, error) {
	return s.repo.FindReplies(ctx, parentID)
}

// GetCommentsByUser returns comments made by a user
func (s *commentServiceImpl) GetCommentsByUser(ctx context.Context, userID primitive.ObjectID) ([]*models.Comment, error) {
	return s.repo.FindByUserID(ctx, userID)
}
