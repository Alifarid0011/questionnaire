package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserAnswerDTO struct {
	QuizID  primitive.ObjectID     `json:"quiz_id" binding:"required"`
	UserID  primitive.ObjectID     `json:"user_id" binding:"required"`
	Answers map[string]interface{} `json:"answers" binding:"required"`
}
