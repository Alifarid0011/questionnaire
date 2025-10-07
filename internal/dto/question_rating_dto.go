package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// QuestionRatingDTO is used for transferring question rating data between layers
type QuestionRatingDTO struct {
	ID         primitive.ObjectID `json:"id,omitempty" swaggerignore:"true"`         // Unique identifier of the rating
	QuestionID primitive.ObjectID `json:"question_id"`                               // ID of the question being rated
	UserID     primitive.ObjectID `json:"user_id" swaggerignore:"true"`              // ID of the user who rated the question
	Score      int                `json:"score"`                                     // Rating score from 1 to 5
	CreatedAt  time.Time          `json:"created_at,omitempty" swaggerignore:"true"` // Timestamp of creation
	UpdatedAt  time.Time          `json:"updated_at,omitempty" swaggerignore:"true"` // Timestamp of last update
}

// UpdateQuestionRatingDTO is used for updating an existing question rating
type UpdateQuestionRatingDTO struct {
	ID    primitive.ObjectID `json:"id" `             // Unique identifier of the rating
	Score *int               `json:"score,omitempty"` // Optional new score to update
}
