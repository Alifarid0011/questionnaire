package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserAnswer struct {
	ID        primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	QuizID    primitive.ObjectID     `bson:"quiz_id" json:"quiz_id"`
	UserID    primitive.ObjectID     `bson:"user_id" json:"user_id"`
	Answers   map[string]interface{} `bson:"answers" json:"answers"` // key: QuestionID, value: answer
	Score     float64                `bson:"score" json:"score"`
	CreatedAt time.Time              `bson:"created_at" json:"created_at"`
}
