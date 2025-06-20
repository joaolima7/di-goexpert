package product

import "database/sql"

type ProductRepositoryInterface interface {
	GetProduct(id int) (Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetProduct(id int) (Product, error) {
	return Product{
		ID:   id,
		Name: "ProductName",
	}, nil
}
