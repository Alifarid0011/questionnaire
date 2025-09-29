package repository

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type questionRatingRepositoryImpl struct {
	collection *mongo.Collection
}

func NewQuestionRatingRepository(db *mongo.Database) QuestionRatingRepository {
	return &questionRatingRepositoryImpl{
		collection: db.Collection("question_ratings"),
	}
}

func (r *questionRatingRepositoryImpl) Create(ctx context.Context, rating *models.QuestionRating) error {
	rating.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, rating)
	return err
}

func (r *questionRatingRepositoryImpl) Update(ctx context.Context, rating *models.QuestionRating) error {
	rating.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": rating.ID}, bson.M{"$set": rating})
	return err
}

func (r *questionRatingRepositoryImpl) FindByID(ctx context.Context, id primitive.ObjectID) (*models.QuestionRating, error) {
	var rating models.QuestionRating
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&rating)
	if err != nil {
		return nil, err
	}
	return &rating, nil
}

func (r *questionRatingRepositoryImpl) FindByQuestionID(ctx context.Context, questionID primitive.ObjectID) ([]*models.QuestionRating, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"question_id": questionID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var ratings []*models.QuestionRating
	if err := cursor.All(ctx, &ratings); err != nil {
		return nil, err
	}
	return ratings, nil
}

func (r *questionRatingRepositoryImpl) FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]*models.QuestionRating, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var ratings []*models.QuestionRating
	if err := cursor.All(ctx, &ratings); err != nil {
		return nil, err
	}
	return ratings, nil
}

func (r *questionRatingRepositoryImpl) FindByQuestionAndUser(ctx context.Context, questionID, userID primitive.ObjectID) (*models.QuestionRating, error) {
	var rating models.QuestionRating
	err := r.collection.FindOne(ctx, bson.M{"question_id": questionID, "user_id": userID}).Decode(&rating)
	if err != nil {
		return nil, err
	}
	return &rating, nil
}

func (r *questionRatingRepositoryImpl) EnsureIndexes(ctx context.Context) error {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "question_id", Value: 1}}, Options: options.Index().SetName("idx_question_id")},
		{Keys: bson.D{{Key: "user_id", Value: 1}}, Options: options.Index().SetName("idx_user_id")},
		{Keys: bson.D{{Key: "question_id", Value: 1}, {Key: "user_id", Value: 1}}, Options: options.Index().SetUnique(true).SetName("idx_question_user")},
	}
	_, err := r.collection.Indexes().CreateMany(ctx, indexes)
	return err
}
