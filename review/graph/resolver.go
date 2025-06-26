package graph

import "github.com/xamenyap/graphql-federation/review/reviews"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ReviewsGetter ReviewsGetter
}

type ReviewsGetter interface {
	GetByProductID(productID string) ([]reviews.Review, error)
}
