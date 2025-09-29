package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAnswerDTO struct {
	QuizID  primitive.ObjectID `json:"quiz_id" binding:"required"`
	UserID  primitive.ObjectID `json:"user_id" binding:"required"`
	Answers []AnswerDTO        `json:"answers" binding:"required"`
}

type AnswerDTO struct {
	QuestionID string   `json:"question_id" binding:"required"`
	Response   []string `json:"response" binding:"required"`
}
