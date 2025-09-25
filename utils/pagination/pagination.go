package pagination

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Pagination interface {
	GetLimit() int
	ApplyToMongoQuery(filter bson.M) (bson.M, *options.FindOptions)
	Count(ctx context.Context, coll *mongo.Collection, filter bson.M) (int64, error)
}

func FromContext(ctx context.Context) (Pagination, bool) {
	p, ok := ctx.Value(constants.PaginatorCtxKey).(Pagination)
	return p, ok
}
