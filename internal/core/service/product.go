package service

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	log "github.com/sirupsen/logrus"
)

type ProductService struct {
	repo    port.IProductRepository
	storage port.IObjectStorage
}

func NewProductService(repo port.IProductRepository, storage port.IObjectStorage) *ProductService {
	return &ProductService{
		repo:    repo,
		storage: storage,
	}
}

func (p *ProductService) Get(ctx context.Context, limit, offset int64) ([]domain.Product, error) {
	product, err := p.repo.Get(ctx, limit, offset)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ProductServiceGet",
		}).Error(err.Error())
		return nil, err
	}

	return product, nil
}

func (p *ProductService) GetByName(ctx context.Context, name string) ([]domain.Product, error) {
	product, err := p.repo.GetByName(ctx, name)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ProductServiceGet",
		}).Error(err.Error())
		return nil, err
	}

	return product, nil
}

func (p *ProductService) GetByID(ctx context.Context, productID domain.ID) (domain.Product, error) {
	product, err := p.repo.GetByID(ctx, productID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ProductServiceGetByID",
		}).Error(err.Error())
		return domain.Product{}, err
	}

	return product, nil
}

func (p *ProductService) Create(ctx context.Context, param port.CreateProductParam) (domain.Product, error) {
	productID := domain.NewID()
	url, err := p.storage.SaveFile(ctx, domain.File{
		Name:   productID.String() + ".png",
		Path:   "product",
		Reader: param.PhotoReader,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ProductServiceCreate",
		}).Error(err.Error())
		return domain.Product{}, err
	}

	product, err := p.repo.Create(ctx, domain.Product{
		ID:          productID,
		Name:        param.Name,
		Description: param.Description,
		Price:       param.Price,
		Category:    param.Category,
		PhotoUrl:    url.String(),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ProductServiceCreate",
		}).Error(err.Error())
		return domain.Product{}, err
	}

	return product, nil
}

func (p *ProductService) Update(ctx context.Context, productID domain.ID, param port.UpdateProductParam) (domain.Product, error) {
	product, err := p.repo.GetByID(ctx, productID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ProductServiceUpdate",
		}).Error(err.Error())
		return domain.Product{}, err
	}

	if param.Name.Valid {
		product.Name = param.Name.String
	}
	if param.Description.Valid {
		product.Description = param.Description.String
	}
	if param.Price.Valid {
		product.Price = param.Price.Int64
	}
	if param.Category != nil {
		product.Category = *param.Category
	}
	if param.PhotoReader != nil {
		url, err := p.storage.SaveFile(ctx, domain.File{
			Name:   productID.String() + ".png",
			Path:   "product",
			Reader: *param.PhotoReader,
		})
		if err != nil {
			log.WithFields(log.Fields{
				"from": "ProductServiceUpdate",
			}).Error(err.Error())
			return domain.Product{}, err
		}
		product.PhotoUrl = url.String()
	}

	product, err = p.repo.Update(ctx, product)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ProductServiceUpdate",
		}).Error(err.Error())
		return domain.Product{}, err
	}

	return product, nil
}

func (p *ProductService) Delete(ctx context.Context, productID domain.ID) error {
	err := p.repo.Delete(ctx, productID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "ProductServiceDelete",
		}).Error(err.Error())
		return err
	}

	return nil
}
