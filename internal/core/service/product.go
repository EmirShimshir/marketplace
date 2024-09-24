package service

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	log "github.com/sirupsen/logrus"
)

type ProductService struct {
	productRepo port.IProductRepository
	storage     port.IObjectStorage
}

func NewProductService(repo port.IProductRepository, storage port.IObjectStorage) *ProductService {
	return &ProductService{
		productRepo: repo,
		storage:     storage,
	}
}

func (p *ProductService) GetByID(ctx context.Context, productID domain.ID) (domain.Product, error) {
	product, err := p.productRepo.GetByID(ctx, productID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ProductServiceGetByID",
		}).Error(err.Error())
		return domain.Product{}, err
	}

	return product, nil
}
