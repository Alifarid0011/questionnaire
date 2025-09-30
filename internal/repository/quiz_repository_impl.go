package repository

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/constant"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type quizRepositoryImpl struct {
	collection *mongo.Collection
}

func NewQuizRepository(db *mongo.Database) QuizRepository {
	return &quizRepositoryImpl{
		collection: db.Collection(constant.QuizCollection),
	}
}

func (r *quizRepositoryImpl) QuizCreate(ctx context.Context, quiz *models.Quiz) error {
	quiz.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, quiz)
	return err
}

func (r *quizRepositoryImpl) QuizFindByID(ctx context.Context, id primitive.ObjectID) (*models.Quiz, error) {
	var quiz models.Quiz
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&quiz)
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (r *quizRepositoryImpl) QuizUpdate(ctx context.Context, quiz *models.Quiz) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": quiz.ID}, bson.M{"$set": quiz})
	return err
}

func (r *quizRepositoryImpl) QuizDelete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *quizRepositoryImpl) QuizGetAll(ctx context.Context) ([]*models.Quiz, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var quizzes []*models.Quiz
	if err := cursor.All(ctx, &quizzes); err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (r *quizRepositoryImpl) QuizGetByCategory(ctx context.Context, category string) ([]*models.Quiz, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"category": category})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var quizzes []*models.Quiz
	if err := cursor.All(ctx, &quizzes); err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (r *quizRepositoryImpl) QuizGetCategories(ctx context.Context) ([]string, error) {
	categories, err := r.collection.Distinct(ctx, "category", bson.D{})
	if err != nil {
		return nil, err
	}
	var result []string
	for _, c := range categories {
		if str, ok := c.(string); ok {
			result = append(result, str)
		}
	}
	return result, nil
}

func (r *quizRepositoryImpl) QuizCountByCategory(ctx context.Context) (map[string]int64, error) {
	pipeline := mongo.Pipeline{
		{{"$group", bson.D{
			{"_id", "$category"},
			{"count", bson.D{{"$sum", 1}}},
		}}},
	}
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	type categoryCount struct {
		ID    string `bson:"_id"`
		Count int64  `bson:"count"`
	}
	var results []categoryCount
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	counts := make(map[string]int64)
	for _, r := range results {
		counts[r.ID] = r.Count
	}
	return counts, nil
}

func (r *quizRepositoryImpl) EnsureIndexes(ctx context.Context) error {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "category", Value: 1}}, Options: options.Index().SetName("idx_category")},
		{Keys: bson.D{{Key: "user_id", Value: 1}}, Options: options.Index().SetName("idx_user_id")},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index().SetName("idx_created_at")},
	}
	_, err := r.collection.Indexes().CreateMany(ctx, indexes)
	return err
}
