package dto

import (
	"github.com/Alifarid0011/questionnaire-back-end/config"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CommentDto struct {
	Id         primitive.ObjectID  `bson:"_id" swaggerignore:"true" json:"id,omitempty"`
	EntityId   primitive.ObjectID  `bson:"entity_id,omitempty" json:"entity_id,omitempty" ` // polymorphic reference
	EntityType string              `bson:"entity_type" json:"entity_type" `                 //
	UserID     primitive.ObjectID  `bson:"user_id,omitempty" json:"user_id,omitempty" swaggerignore:"true"`
	Text       string              `bson:"text" json:"text"`
	ParentID   *primitive.ObjectID `bson:"parent_id,omitempty" json:"parent_id,omitempty"`
}

func (d *CommentDto) ToModel(userID primitive.ObjectID, cratedAt *time.Time) *models.Comment {
	return &models.Comment{
		ID:        primitive.NewObjectID(),
		Target:    models.DBRef{Ref: d.EntityType, ID: d.EntityId, DB: config.Get.Mongo.DbName},
		UserID:    userID,
		Text:      d.Text,
		ParentID:  d.ParentID,
		CreatedAt: cratedAt,
		UpdatedAt: time.Now(),
	}
}
