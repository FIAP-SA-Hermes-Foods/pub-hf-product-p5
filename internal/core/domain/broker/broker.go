package broker

import (
	"pub-hf-product-p5/internal/core/domain/entity/dto"
)

type ProductBroker interface {
	GetProductByID(input dto.ProductBroker) error
	SaveProduct(input dto.ProductBroker) error
	UpdateProductByID(input dto.ProductBroker) error
	GetProductByCategory(input dto.ProductBroker) error
	DeleteProductByID(input dto.ProductBroker) error
}
