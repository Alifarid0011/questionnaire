package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type QuestionRating struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	QuestionID primitive.ObjectID `bson:"question_id" json:"question_id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	Score      int                `bson:"score" json:"score"` // 1  5
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
