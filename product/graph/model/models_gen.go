// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Product struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Price  float64  `json:"price"`
	Weight *float64 `json:"weight,omitempty"`
}

func (Product) IsEntity() {}

type Query struct {
}
