package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Comment struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Target    DBRef               `bson:"target" json:"target"` // polymorphic reference
	UserID    primitive.ObjectID  `bson:"user_id" json:"user_id"`
	Text      string              `bson:"text" json:"text"`
	ParentID  *primitive.ObjectID `bson:"parent_id,omitempty" json:"parent_id,omitempty"`
	CreatedAt time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
