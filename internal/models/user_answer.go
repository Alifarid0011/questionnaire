package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserAnswer struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	QuizID    primitive.ObjectID `bson:"quiz_id" json:"quiz_id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Answers   []Answer           `bson:"answers" json:"answers"`
	Score     float64            `bson:"score" json:"score"`
	Appeal    bool               `bson:"appeal" json:"appeal"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type Answer struct {
	QuestionID string   `bson:"question_id" json:"question_id"`
	Response   []string `bson:"response" json:"response"` // generic: short answer = [text], radio = [choice], checkbox = [choices]
	Score      float64  `bson:"score" json:"score"`
	IsCorrect  bool     `bson:"is_correct" json:"is_correct"`
	Comment    string   `bson:"comment,omitempty" json:"comment,omitempty"`
}
