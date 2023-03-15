package usecase

import (
	"edmar.lima/edmarlima/product-api/internal/entity"
)

type ListProductsOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type ListProductsUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductsUseCase(productRepository entity.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{ProductRepository: productRepository}
}

func (u *ListProductsUseCase) Execute() ([]*ListProductsOutputDto, error) {

	products, err := u.ProductRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var productsOuput []*ListProductsOutputDto

	for _, product := range products {
		productsOuput = append(productsOuput, &ListProductsOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return productsOuput, nil
}
