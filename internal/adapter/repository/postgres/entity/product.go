package entity

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/google/uuid"
)

const (
	PgProductElectronic = "Electronic"
	PgProductFashion    = "Fashion"
	PgProductHome       = "Home"
	PgProductHealth     = "Health"
	PgProductSport      = "Sport"
	PgProductBooks      = "Books"
)

type PgProduct struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       int64     `db:"price"`
	Category    string    `db:"category"`
	PhotoUrl    string    `db:"photo_url"`
}

func (u *PgProduct) ToDomain() domain.Product {
	var productCategory domain.ProductCategory
	switch u.Category {
	case PgProductElectronic:
		productCategory = domain.ElectronicCategory
	case PgProductFashion:
		productCategory = domain.FashionCategory
	case PgProductHome:
		productCategory = domain.HomeCategory
	case PgProductHealth:
		productCategory = domain.HealthCategory
	case PgProductSport:
		productCategory = domain.SportCategory
	case PgProductBooks:
		productCategory = domain.BooksCategory
	}
	return domain.Product{
		ID:          domain.ID(u.ID.String()),
		Name:        u.Name,
		Description: u.Description,
		Price:       u.Price,
		Category:    productCategory,
		PhotoUrl:    u.PhotoUrl,
	}
}

func NewPgProduct(product domain.Product) PgProduct {
	id, _ := uuid.Parse(product.ID.String())
	var productCategory string
	switch product.Category {
	case domain.ElectronicCategory:
		productCategory = PgProductElectronic
	case domain.FashionCategory:
		productCategory = PgProductFashion
	case domain.HomeCategory:
		productCategory = PgProductHome
	case domain.HealthCategory:
		productCategory = PgProductHealth
	case domain.SportCategory:
		productCategory = PgProductSport
	case domain.BooksCategory:
		productCategory = PgProductBooks
	}
	return PgProduct{
		ID:          id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    productCategory,
		PhotoUrl:    product.PhotoUrl,
	}
}
