package repository

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/constant"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type commentRepositoryImpl struct {
	collection *mongo.Collection
}

// NewCommentRepository creates a new comment repository
func NewCommentRepository(db *mongo.Database) CommentRepository {
	return &commentRepositoryImpl{
		collection: db.Collection(constant.CommentCollection),
	}
}

// Create inserts a new comment
func (r *commentRepositoryImpl) Create(ctx context.Context, comment *models.Comment) error {
	_, err := r.collection.InsertOne(ctx, comment)
	return err
}

// Update modifies an existing comment
func (r *commentRepositoryImpl) Update(ctx context.Context, comment *models.Comment) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": comment.ID}, bson.M{"$set": comment})
	return err
}

// FindByID retrieves a comment by its ID
func (r *commentRepositoryImpl) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Comment, error) {
	var comment models.Comment
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// FindByTarget retrieves comments by polymorphic target (DBRef)
func (r *commentRepositoryImpl) FindByTarget(ctx context.Context, target models.DBRef) ([]*models.Comment, error) {
	filter := bson.M{
		"target.$ref": target.Ref,
		"target.$id":  target.ID,
	}
	cursor, err := r.collection.Find(ctx, filter)
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

// FindByUserID retrieves comments by user
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

// FindReplies retrieves comments that are replies to a parent comment
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

// EnsureIndexes creates indexes for faster queries
func (r *commentRepositoryImpl) EnsureIndexes(ctx context.Context) error {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "target.$ref", Value: 1}}, Options: options.Index().SetName("idx_target_ref")},
		{Keys: bson.D{{Key: "target.$id", Value: 1}}, Options: options.Index().SetName("idx_target_id")},
		{Keys: bson.D{{Key: "user_id", Value: 1}}, Options: options.Index().SetName("idx_user_id")},
		{Keys: bson.D{{Key: "parent_id", Value: 1}}, Options: options.Index().SetName("idx_parent_id")},
	}
	_, err := r.collection.Indexes().CreateMany(ctx, indexes)
	return err
}
