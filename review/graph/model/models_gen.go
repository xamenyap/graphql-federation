// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Product struct {
	ID      string    `json:"id"`
	Reviews []*Review `json:"reviews,omitempty"`
}

func (Product) IsEntity() {}

type Query struct {
}

type Review struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}
