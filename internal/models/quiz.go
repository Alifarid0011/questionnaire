package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Quiz struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Category  string             `bson:"category" json:"category"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Level     string             `json:"level" bson:"level"`
	Questions []Question         `bson:"questions" json:"questions"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type Question struct {
	ID            string   `bson:"id" json:"id"`
	Type          string   `bson:"type" json:"type"` // short, checkbox, radio
	Label         string   `bson:"label" json:"label"`
	Options       []string `bson:"options" json:"options"`
	CorrectAnswer []string `bson:"correct_answer" json:"correct_answer"`
	KeyWords      []string `json:"key_words" bson:"key_words"`
}
