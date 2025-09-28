package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateUserAnswerDTO struct {
	QuizID  primitive.ObjectID     `json:"quiz_id" binding:"required"`
	UserID  primitive.ObjectID     `json:"user_id" binding:"required"`
	Answers map[string]interface{} `json:"answers" binding:"required"`
}

type UpdateUserAnswerDTO struct {
	ID      primitive.ObjectID     `json:"id" binding:"required"`
	Answers map[string]interface{} `json:"answers,omitempty"`
	Score   *float64               `json:"score,omitempty"`
}
