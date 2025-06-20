package product

type ProductUseCase struct {
	repository ProductRepository
}

func NewProductUseCase(repository ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repository: repository,
	}
}

func (pu *ProductUseCase) GetProduct(id int) (Product, error) {
	return pu.repository.GetProductById(id)
}
