package service

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
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
func (s *commentServiceImpl) CreateComment(ctx context.Context, comment *dto.CommentDto) (*dto.CommentDto, error) {
	CreatedAt := time.Now()
	if err := s.repo.Create(ctx, comment.ToModel(ctx.Value("user_uid").(primitive.ObjectID), &CreatedAt)); err != nil {
		return nil, err
	}
	return comment, nil
}

// UpdateComment updates an existing comment
func (s *commentServiceImpl) UpdateComment(ctx context.Context, comment *dto.CommentDto) (*dto.CommentDto, error) {
	if err := s.repo.Update(ctx, comment.ToModel(ctx.Value("user_uid").(primitive.ObjectID), nil)); err != nil {
		return nil, err
	}
	return comment, nil
}

// GetCommentByID returns a comment by its ID
func (s *commentServiceImpl) GetCommentByID(ctx context.Context, id primitive.ObjectID) (*dto.CommentDto, error) {
	comment, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.CommentDto{
		Id:         comment.ID,
		ParentID:   comment.ParentID,
		EntityType: comment.Target.Ref,
		UserID:     comment.UserID,
		Text:       comment.Text,
	}, nil
}

// GetCommentsByTarget returns comments for a polymorphic target
func (s *commentServiceImpl) GetCommentsByTarget(ctx context.Context, target models.DBRef) ([]*dto.CommentDto, error) {
	comments, err := s.repo.FindByTarget(ctx, target)
	if err != nil {
		return nil, err
	}
	var commentDto []*dto.CommentDto
	for _, comment := range comments {
		commentDto = append(commentDto, &dto.CommentDto{
			Id:         comment.ID,
			ParentID:   comment.ParentID,
			EntityType: comment.Target.Ref,
			EntityId:   comment.Target.ID,
			UserID:     comment.UserID,
			Text:       comment.Text,
		})
	}
	return commentDto, nil
}

// GetReplies returns replies to a parent comment
func (s *commentServiceImpl) GetReplies(ctx context.Context, parentID primitive.ObjectID) ([]*dto.CommentDto, error) {
	comments, err := s.repo.FindReplies(ctx, parentID)
	if err != nil {
		return nil, err
	}
	var commentDto []*dto.CommentDto
	for _, comment := range comments {
		commentDto = append(commentDto, &dto.CommentDto{
			Id:       comment.ID,
			ParentID: comment.ParentID,
			EntityId: comment.Target.ID,

			EntityType: comment.Target.Ref,
			UserID:     comment.UserID,
			Text:       comment.Text,
		})
	}
	return commentDto, nil
}

// GetCommentsByUser returns comments made by a user
func (s *commentServiceImpl) GetCommentsByUser(ctx context.Context, userID primitive.ObjectID) ([]*dto.CommentDto, error) {
	comments, err := s.repo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var commentDto []*dto.CommentDto
	for _, comment := range comments {
		commentDto = append(commentDto, &dto.CommentDto{
			Id:         comment.ID,
			ParentID:   comment.ParentID,
			EntityId:   comment.Target.ID,
			EntityType: comment.Target.Ref,
			UserID:     comment.UserID,
			Text:       comment.Text,
		})
	}
	return commentDto, nil
}
