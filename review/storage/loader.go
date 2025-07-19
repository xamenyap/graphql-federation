package storage

import (
	"context"
	"time"

	"github.com/vikstrous/dataloadgen"
)

func NewLoader(repo *Repository) *dataloadgen.Loader[string, []Review] {
	r := &reader{
		repo: repo,
	}
	return dataloadgen.NewLoader(r.getByProductIDs, dataloadgen.WithWait(time.Millisecond))
}

type reader struct {
	repo *Repository
}

func (r *reader) getByProductIDs(ctx context.Context, productIDs []string) ([][]Review, []error) {
	results, err := r.repo.GetByProductIDs(ctx, productIDs)
	if err != nil {
		return nil, []error{err}
	}

	reviews := make([][]Review, 0)
	for _, productID := range productIDs {
		reviews = append(reviews, results[productID])
	}

	return reviews, nil
}
