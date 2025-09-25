package pagination

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PageDriver struct {
	Page     int
	PerPage  int
	SortBy   string
	AscOrder bool
	Total    int64
}

func NewPageDriver(page, perPage int, sortBy string, asc bool) *PageDriver {
	return &PageDriver{
		Page:     page,
		PerPage:  perPage,
		SortBy:   sortBy,
		AscOrder: asc,
	}
}

func (p *PageDriver) GetLimit() int {
	return p.PerPage
}

func (p *PageDriver) ApplyToMongoQuery(filter bson.M) (bson.M, *options.FindOptions) {
	if filter == nil {
		filter = bson.M{}
	}
	skip := (p.Page - 1) * p.PerPage
	order := -1
	if p.AscOrder {
		order = 1
	}
	opts := options.Find().
		SetLimit(int64(p.PerPage)).
		SetSkip(int64(skip)).
		SetSort(bson.D{{Key: p.SortBy, Value: order}})

	return filter, opts
}

func (p *PageDriver) Count(ctx context.Context, coll *mongo.Collection, filter bson.M) (int64, error) {
	if filter == nil {
		filter = bson.M{}
	}
	total, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	p.Total = total
	return total, nil
}
