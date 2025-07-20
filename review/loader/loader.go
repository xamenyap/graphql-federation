package loader

import (
	"context"
	"net/http"
	"time"

	"github.com/vikstrous/dataloadgen"
	"github.com/xamenyap/graphql-federation/review/storage"
)

func NewLoader(repo *storage.Repository) *dataloadgen.Loader[string, []storage.Review] {
	r := &reader{
		repo: repo,
	}
	return dataloadgen.NewLoader(r.getByProductIDs, dataloadgen.WithWait(time.Millisecond))
}

type reader struct {
	repo *storage.Repository
}

func (r *reader) getByProductIDs(ctx context.Context, productIDs []string) ([][]storage.Review, []error) {
	results, err := r.repo.GetByProductIDs(ctx, productIDs)
	if err != nil {
		return nil, []error{err}
	}

	reviews := make([][]storage.Review, 0)
	for _, productID := range productIDs {
		reviews = append(reviews, results[productID])
	}

	return reviews, nil
}

type ctxKey string

const loaderCtxKey ctxKey = "loader"

// Middleware injects data loaders into the context
func Middleware(repo *storage.Repository, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := NewLoader(repo)
		r = r.WithContext(context.WithValue(r.Context(), loaderCtxKey, l))
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *dataloadgen.Loader[string, []storage.Review] {
	return ctx.Value(loaderCtxKey).(*dataloadgen.Loader[string, []storage.Review])
}
