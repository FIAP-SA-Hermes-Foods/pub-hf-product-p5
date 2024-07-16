package dto

import (
	"pub-hf-product-p5/internal/core/domain/entity"
	vo "pub-hf-product-p5/internal/core/domain/entity/valueObject"
)

type ProductBroker struct {
	UUID          string  `json:"uuid,omitempty"`
	MessageID     string  `json:"messageId,omitempty"`
	Name          string  `json:"name,omitempty"`
	Category      string  `json:"category,omitempty"`
	Image         string  `json:"image,omitempty"`
	Description   string  `json:"description,omitempty"`
	Price         float64 `json:"price,omitempty"`
	CreatedAt     string  `json:"createdAt,omitempty"`
	DeactivatedAt string  `json:"deactivatedAt,omitempty"`
}

type (
	RequestProduct struct {
		UUID          string  `json:"uuid,omitempty"`
		Name          string  `json:"name,omitempty"`
		Category      string  `json:"category,omitempty"`
		Image         string  `json:"image,omitempty"`
		Description   string  `json:"description,omitempty"`
		Price         float64 `json:"price,omitempty"`
		CreatedAt     string  `json:"createdAt,omitempty"`
		DeactivatedAt string  `json:"deactivatedAt,omitempty"`
	}

	OutputProduct struct {
		UUID          string  `json:"uuid,omitempty"`
		Name          string  `json:"name,omitempty"`
		Category      string  `json:"category,omitempty"`
		Image         string  `json:"image,omitempty"`
		Description   string  `json:"description,omitempty"`
		Price         float64 `json:"price,omitempty"`
		CreatedAt     string  `json:"createdAt,omitempty"`
		DeactivatedAt string  `json:"deactivatedAt,omitempty"`
	}
)

func (r RequestProduct) Product() entity.Product {
	product := entity.Product{
		Name: r.Name,
		Category: vo.Category{
			Value: r.Category,
		},
		Image:       r.Image,
		Description: r.Description,
		Price:       r.Price,
	}

	return product
}
