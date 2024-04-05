package repository

import (
	"app/internal"
	"database/sql"
	"errors"
)

var (
	// ErrProductRowsAffected is an error that is returned when the number of rows affected is not 1
	ErrProductRowsAffected = errors.New("number of rows affected is not the expected")
)

// NewProductMySQL returns a new instance of ProductMySQL
func NewProductMySQL(db *sql.DB) *ProductMySQL {
	return &ProductMySQL{
		db: db,
	}
}

// ProductMySQL is a struct that represents a MySQL implementation of the ProductRepository
type ProductMySQL struct {
	// db is a pool of database connections
	db *sql.DB
}

// FindById finds a product by id.
func (r *ProductMySQL) FindById(id int) (p internal.Product, err error) {
	// execute query
	query := "SELECT id, name, quantity, code_value, is_published,expiration, price FROM products WHERE id = ?"
	row := r.db.QueryRow(query, id)
	if err = row.Err(); err != nil {
		return
	}

	// scan the result
	err = row.Scan(&p.Id, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrRepositoryProductNotFound
			return
		}
		return
	}

	return
}

// Save saves a product.
func (r *ProductMySQL) Save(p *internal.Product) (err error) {
	// execute the query
	result, err := r.db.Exec(
		"INSERT INTO `products` (`name`, `quantity`, `code_value`, `is_published`, `expiration`, `price`) "+
			"VALUES (?,?,?,?,?,?)",
		p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price,
	)
	if err != nil {
		return
	}

	// get the last inserted id
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	p.Id = int(id)

	return
}

// UpdateOrSave updates or saves a product.
func (r *ProductMySQL) UpdateOrSave(p *internal.Product) (err error) {
	return
}

// Update updates a product.
func (r *ProductMySQL) Update(p *internal.Product) (err error) {
	return
}

// Delete deletes a product.
func (r *ProductMySQL) Delete(id int) (err error) {
	return
}
