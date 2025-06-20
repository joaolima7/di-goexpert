//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/joaolima7/di-goexpert/product"
)

var setRepositoryDependence = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewProductUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependence,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
