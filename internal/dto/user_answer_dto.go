package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAnswerDTO struct {
	ID      primitive.ObjectID  `bson:"_id,omitempty" json:"id" `
	QuizID  primitive.ObjectID  `json:"quiz_id" binding:"required"`
	UserID  *primitive.ObjectID `json:"user_id"`
	Answers []AnswerDTO         `json:"answers" binding:"required"`
}

type AnswerDTO struct {
	QuestionID primitive.ObjectID `json:"question_id" binding:"required"`
	Response   []string           `json:"response" binding:"required"`
}
