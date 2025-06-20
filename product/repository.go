package product

import "database/sql"

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetProductById(id int) (Product, error) {
	return Product{
		ID:   id,
		Name: "ProductName",
	}, nil
}
