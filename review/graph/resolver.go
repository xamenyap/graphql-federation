package graph

import (
	"context"

	"github.com/xamenyap/graphql-federation/review/storage"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ReviewsGetter ReviewsGetter
}

type ReviewsGetter interface {
	GetByProductID(ctx context.Context, productID string) ([]storage.Review, error)
}
