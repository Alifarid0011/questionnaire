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

type commentRepositoryImpl struct {
	collection *mongo.Collection
}

func NewCommentRepository(db *mongo.Database) CommentRepository {
	return &commentRepositoryImpl{
		collection: db.Collection("comments"),
	}
}

func (r *commentRepositoryImpl) Create(ctx context.Context, comment *models.Comment) error {
	comment.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, comment)
	return err
}

func (r *commentRepositoryImpl) Update(ctx context.Context, comment *models.Comment) error {
	comment.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": comment.ID}, bson.M{"$set": comment})
	return err
}

func (r *commentRepositoryImpl) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Comment, error) {
	var comment models.Comment
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepositoryImpl) FindByQuestionID(ctx context.Context, questionID primitive.ObjectID) ([]*models.Comment, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"question_id": questionID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []*models.Comment
	if err := cursor.All(ctx, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepositoryImpl) FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]*models.Comment, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []*models.Comment
	if err := cursor.All(ctx, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepositoryImpl) FindReplies(ctx context.Context, parentID primitive.ObjectID) ([]*models.Comment, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"parent_id": parentID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var replies []*models.Comment
	if err := cursor.All(ctx, &replies); err != nil {
		return nil, err
	}
	return replies, nil
}

func (r *commentRepositoryImpl) EnsureIndexes(ctx context.Context) error {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "question_id", Value: 1}}, Options: options.Index().SetName("idx_question_id")},
		{Keys: bson.D{{Key: "user_id", Value: 1}}, Options: options.Index().SetName("idx_user_id")},
		{Keys: bson.D{{Key: "parent_id", Value: 1}}, Options: options.Index().SetName("idx_parent_id")},
	}
	_, err := r.collection.Indexes().CreateMany(ctx, indexes)
	return err
}
