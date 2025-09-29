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

type userAnswerRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserAnswerRepository(db *mongo.Database) UserAnswerRepository {
	return &userAnswerRepositoryImpl{
		collection: db.Collection("user_answers"),
	}
}

func (r *userAnswerRepositoryImpl) UserAnswerCreate(ctx context.Context, answer *models.UserAnswer) error {
	answer.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, answer)
	return err
}

func (r *userAnswerRepositoryImpl) UserAnswerFindByID(ctx context.Context, id primitive.ObjectID) (*models.UserAnswer, error) {
	var answer models.UserAnswer
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&answer)
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

func (r *userAnswerRepositoryImpl) UserAnswerFindByQuizID(ctx context.Context, quizID primitive.ObjectID) ([]*models.UserAnswer, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"quiz_id": quizID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var answers []*models.UserAnswer
	if err := cursor.All(ctx, &answers); err != nil {
		return nil, err
	}
	return answers, nil
}

func (r *userAnswerRepositoryImpl) UserAnswerFindByUserID(ctx context.Context, userID primitive.ObjectID) ([]*models.UserAnswer, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var answers []*models.UserAnswer
	if err := cursor.All(ctx, &answers); err != nil {
		return nil, err
	}
	return answers, nil
}

func (r *userAnswerRepositoryImpl) UserAnswerFindByQuizIDAndUserID(ctx context.Context, quizID, userID primitive.ObjectID) ([]*models.UserAnswer, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"quiz_id": quizID, "user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var answers []*models.UserAnswer
	if err := cursor.All(ctx, &answers); err != nil {
		return nil, err
	}
	return answers, nil
}

func (r *userAnswerRepositoryImpl) UserAnswerEnsureIndexes(ctx context.Context) error {
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "quiz_id", Value: 1}},
			Options: options.Index().SetName("idx_quiz_id"),
		},
		{
			Keys:    bson.D{{Key: "user_id", Value: 1}},
			Options: options.Index().SetName("idx_user_id"),
		},
		{
			Keys:    bson.D{{Key: "quiz_id", Value: 1}, {Key: "user_id", Value: 1}},
			Options: options.Index().SetName("idx_quiz_user"),
		},
	}
	_, err := r.collection.Indexes().CreateMany(ctx, indexes)
	return err
}
