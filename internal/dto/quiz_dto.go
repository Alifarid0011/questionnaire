package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type QuizDTO struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id" swaggerignore:"true"`
	Title     string              `json:"title" binding:"required"`
	Category  string              `json:"category" binding:"required"`
	Level     string              `json:"level" binding:"required"`
	UserID    *primitive.ObjectID `json:"user_id" swaggerignore:"true"`
	Questions []QuestionDTO       `json:"questions" binding:"required,dive"`
	CreatedAt time.Time           `json:"created_at" bson:"created_at"`
}

type UpdateQuizDTO struct {
	ID        primitive.ObjectID `json:"id" binding:"required"`
	Title     *string            `json:"title,omitempty"`
	Category  *string            `json:"category,omitempty"`
	Level     *string            `json:"level,omitempty"`
	Questions *[]QuestionDTO     `json:"questions,omitempty"`
}

type QuestionDTO struct {
	ID            primitive.ObjectID `json:"id" swaggerignore:"true"`
	Type          string             `json:"type" binding:"required,oneof=short checkbox radio"`
	Label         string             `json:"label" binding:"required"`
	Options       []string           `json:"options,omitempty"`
	CorrectAnswer []string           `json:"correct_answer,omitempty"`
	KeyWords      []string           `json:"key_words,omitempty"`
}
