package repository

import (
	"context"
	"errors"
	"github.com/Alifarid0011/questionnaire-back-end/constant"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"github.com/Alifarid0011/questionnaire-back-end/utils/pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type UserRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &UserRepositoryImpl{
		collection: db.Collection(constant.UserCollection),
	}
}

func (r *UserRepositoryImpl) Delete(user *models.User, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.collection.DeleteOne(ctx, bson.M{"uid": user.UID})
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Update(user *models.User, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"uid": user.UID},
		bson.M{
			"$set": user,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) FindByUsername(username string, ctx context.Context) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindByUID(uid primitive.ObjectID, ctx context.Context) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"uid": uid}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Create(user *models.User, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepositoryImpl) GetAll(ctx context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	filter := bson.M{}
	opts := options.Find()
	if paginator, ok := pagination.FromContext(ctx); ok {
		filter, opts = paginator.ApplyToMongoQuery(filter)
		_, err := paginator.Count(ctx, r.collection, filter)
		if err != nil {
			return nil, err
		}
		cursor, err := r.collection.Find(ctx, filter, opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		var users []models.User
		if err := cursor.All(ctx, &users); err != nil {
			return nil, err
		}

		return users, nil
	}
	return nil, errors.New("failed to find all users")

}
func (r *UserRepositoryImpl) EnsureIndexes(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "username", Value: 1}}, Options: options.Index().SetName("idx_username").SetUnique(true)},
		{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetName("idx_email").SetUnique(true)},
		{Keys: bson.D{{Key: "mobile", Value: 1}}, Options: options.Index().SetName("idx_mobile").SetUnique(true)},
		{Keys: bson.D{{Key: "full_name", Value: 1}}, Options: options.Index().SetName("idx_full_name")},
		{Keys: bson.D{{Key: "created_at", Value: 1}}, Options: options.Index().SetName("idx_created_at")},
		{Keys: bson.D{{Key: "updated_at", Value: 1}}, Options: options.Index().SetName("idx_updated_at")},
		{Keys: bson.D{{Key: "uid", Value: 1}}, Options: options.Index().SetName("uid_unique").SetUnique(true)},
	}
	_, err := r.collection.Indexes().CreateMany(ctx, indexes)
	return err
}
