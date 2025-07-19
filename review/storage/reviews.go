package storage

import (
	"database/sql"
	"fmt"
)

type Review struct {
	ID        string
	Body      string
	ProductID string
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByProductID(productID string) ([]Review, error) {
	query := `SELECT id, body, product_id FROM reviews WHERE product_id = ?`

	rows, err := r.db.Query(query, productID)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var r Review
		if err := rows.Scan(&r.ID, &r.Body, &r.ProductID); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		reviews = append(reviews, r)
	}

	return reviews, nil
}

func (r *Repository) Seed() error {
	createTable := `
	CREATE TABLE IF NOT EXISTS reviews (
		id VARCHAR(255) PRIMARY KEY,
		body TEXT,
		product_id VARCHAR(255),
		INDEX idx_product_id (product_id)
	);`

	if _, err := r.db.Exec(createTable); err != nil {
		return fmt.Errorf("create table error: %w", err)
	}

	inserts := []struct {
		id        string
		body      string
		productID int
	}{
		{
			id:        "1",
			body:      "nice product, will buy again",
			productID: 1,
		},
		{
			id:        "2",
			body:      "very decent, nice seller, may consider buying again",
			productID: 1,
		},
		{
			id:        "3",
			body:      "very low quality, should not buy from this seller",
			productID: 2,
		},
	}

	stm := `INSERT INTO reviews (id, body, product_id) VALUES (?, ?, ?);`
	for _, insert := range inserts {
		if _, err := r.db.Exec(stm, insert.id, insert.body, insert.productID); err != nil {
			return fmt.Errorf("insert review failed: %w", err)
		}
	}

	return nil
}
