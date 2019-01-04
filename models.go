package main

import (
	"database/sql"
	"errors"
)

// This unexported struct will define the needed data for the product that we will use to interact with the db
type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}
func (p *product) updateProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}
func (p *product) deleteProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}
func (p *product) createProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

// Get a list of products (this is a standalone func notice it is not tied to a struct, thus making it standalone)
func getProducts(db *sql.DB, start, count int) ([]product, error) {
	return nil, errors.New("Not implemented")
}
