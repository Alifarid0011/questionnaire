package repository

import (
	"context"
	"fmt"
)

type IndexEnsurer interface {
	EnsureIndexes(ctx context.Context) error
}

func EnsureAllIndexes(ctx context.Context, repos []IndexEnsurer) error {
	for _, r := range repos {
		if err := r.EnsureIndexes(ctx); err != nil {
			return fmt.Errorf("failed to ensure indexes: %w", err)
		}
	}
	return nil
}
