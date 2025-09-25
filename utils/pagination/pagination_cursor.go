package pagination

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CursorDriver struct {
	LastSeenID primitive.ObjectID
	PerPage    int
	SortBy     string
	AscOrder   bool
	HasCursor  bool
}

func (c *CursorDriver) Count(ctx context.Context, coll *mongo.Collection, filter bson.M) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func NewCursorDriver(lastSeenID string, perPage int, sortBy string, asc bool) *CursorDriver {
	driver := &CursorDriver{
		PerPage:  perPage,
		SortBy:   sortBy,
		AscOrder: asc,
	}

	if id, err := primitive.ObjectIDFromHex(lastSeenID); err == nil {
		driver.LastSeenID = id
		driver.HasCursor = true
	}

	return driver
}

func (c *CursorDriver) GetLimit() int {
	return c.PerPage
}

func (c *CursorDriver) ApplyToMongoQuery(filter bson.M) (bson.M, *options.FindOptions) {
	if filter == nil {
		filter = bson.M{}
	}

	order := 1
	compareOp := "$gt"
	if !c.AscOrder {
		order = -1
		compareOp = "$lt"
	}

	if c.HasCursor {
		filter["_id"] = bson.M{compareOp: c.LastSeenID}
	}

	opts := options.Find().
		SetLimit(int64(c.PerPage)).
		SetSort(bson.D{{Key: c.SortBy, Value: order}})

	return filter, opts
}
