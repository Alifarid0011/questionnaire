package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBRef struct {
	Ref string             `bson:"$ref" json:"$ref" binding:"required"` // Collection name
	ID  primitive.ObjectID `bson:"$id" json:"$id" binding:"required"`   // Document ID
	DB  string             `bson:"$db,omitempty" json:"$db,omitempty"`  // Optional DB name
}
