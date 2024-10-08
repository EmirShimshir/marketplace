package service

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	log "github.com/sirupsen/logrus"
)

type ShopService struct {
	shopRepo port.IShopRepository
	storage  port.IObjectStorage
}

func NewShopService(repo port.IShopRepository, storage port.IObjectStorage) *ShopService {
	return &ShopService{
		shopRepo: repo,
		storage:  storage,
	}
}

func (s *ShopService) GetShopByID(ctx context.Context, shopID domain.ID) (domain.Shop, error) {
	shop, err := s.shopRepo.GetShopByID(ctx, shopID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetShopByID",
		}).Error(err.Error())
		return domain.Shop{}, err
	}

	log.WithFields(log.Fields{
		"SellerID": shop.SellerID,
	}).Info("GetShopByID OK")
	return shop, nil
}

func (s *ShopService) GetShopBySellerID(ctx context.Context, sellerID domain.ID) ([]domain.Shop, error) {
	shop, err := s.shopRepo.GetShopBySellerID(ctx, sellerID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetShopBySellerID",
		}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{
		"count": len(shop),
	}).Info("GetShopBySellerID OK")
	return shop, nil
}

func (s *ShopService) CreateShop(ctx context.Context, sellerID domain.ID, param port.CreateShopParam) (domain.Shop, error) {
	if param.Name == "" {
		log.WithFields(log.Fields{
			"from": "CreateShop",
		}).Error(domain.ErrName.Error())
		return domain.Shop{}, domain.ErrName
	}
	if param.Description == "" {
		log.WithFields(log.Fields{
			"from": "CreateShop",
		}).Error(domain.ErrDescription.Error())
		return domain.Shop{}, domain.ErrDescription
	}
	if param.Requisites == "" {
		log.WithFields(log.Fields{
			"from": "CreateShop",
		}).Error(domain.ErrRequisites.Error())
		return domain.Shop{}, domain.ErrRequisites
	}

	shop, err := s.shopRepo.CreateShop(ctx, domain.Shop{
		ID:          domain.NewID(),
		SellerID:    sellerID,
		Name:        param.Name,
		Description: param.Description,
		Requisites:  param.Requisites,
		Email:       param.Email,
		Items:       make([]domain.ShopItem, 0),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateShop",
		}).Error(err.Error())
		return domain.Shop{}, err
	}

	log.WithFields(log.Fields{
		"SellerID": shop.SellerID,
		"Email":    shop.Email,
	}).Info("CreateShop OK")
	return shop, nil
}

func (s *ShopService) GetShopItems(ctx context.Context, limit, offset int64) ([]domain.ShopItem, error) {
	shopItems, err := s.shopRepo.GetShopItems(ctx, limit, offset)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetShopItems",
		}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{
		"count": len(shopItems),
	}).Info("GetShopItems OK")
	return shopItems, nil
}

func (s *ShopService) GetShopItemByProductID(ctx context.Context, productID domain.ID) (domain.ShopItem, error) {
	shopItem, err := s.shopRepo.GetShopItemByProductID(ctx, productID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetShopItemByProductID",
		}).Error(err.Error())
		return domain.ShopItem{}, err
	}

	log.WithFields(log.Fields{
		"shopItemID": shopItem.ID,
	}).Info("GetShopItemByProductID OK")
	return shopItem, nil
}

func (s *ShopService) CreateShopItem(ctx context.Context, param port.CreateShopItemParam) (domain.ShopItem, error) {
	productID := domain.NewID()
	url, err := s.storage.SaveFile(ctx, domain.File{
		Name:   productID.String() + ".png",
		Path:   "product",
		Reader: param.ProductParam.PhotoReader,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateShopItem",
		}).Error(err.Error())
		return domain.ShopItem{}, err
	}

	if param.ProductParam.Name == "" {
		log.WithFields(log.Fields{
			"from": "CreateShopItem",
		}).Error(domain.ErrName.Error())
		return domain.ShopItem{}, domain.ErrName
	}
	if param.ProductParam.Description == "" {
		log.WithFields(log.Fields{
			"from": "CreateShopItem",
		}).Error(domain.ErrDescription.Error())
		return domain.ShopItem{}, domain.ErrDescription
	}
	if param.Quantity < 1 {
		log.WithFields(log.Fields{
			"from": "CreateShopItem",
		}).Error(domain.ErrQuantityItems.Error())
		return domain.ShopItem{}, domain.ErrQuantityItems
	}
	if param.ProductParam.Price < 1 {
		log.WithFields(log.Fields{
			"from": "CreateShopItem",
		}).Error(domain.ErrPrice.Error())
		return domain.ShopItem{}, domain.ErrPrice
	}

	product := domain.Product{
		ID:          productID,
		Name:        param.ProductParam.Name,
		Description: param.ProductParam.Description,
		Price:       param.ProductParam.Price,
		Category:    param.ProductParam.Category,
		PhotoUrl:    url.String(),
	}

	shopItem := domain.ShopItem{
		ID:        domain.NewID(),
		ShopID:    param.ShopID,
		ProductID: product.ID,
		Quantity:  param.Quantity,
	}

	shopItem, err = s.shopRepo.CreateShopItem(ctx, shopItem, product)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "CreateShopItem",
		}).Error(err.Error())
		return domain.ShopItem{}, err
	}

	log.WithFields(log.Fields{
		"shopItemID": shopItem.ID,
	}).Info("CreateShopItem OK")
	return shopItem, nil
}

func (s *ShopService) UpdateShopItem(ctx context.Context, shopItemID domain.ID, param port.UpdateShopItemParam) (domain.ShopItem, error) {
	shopItem, err := s.shopRepo.GetShopItemByID(ctx, shopItemID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateShopItem",
		}).Error(err.Error())
		return domain.ShopItem{}, err
	}

	if param.Quantity.Valid {
		if param.Quantity.Int64 < 0 {
			log.WithFields(log.Fields{
				"from": "UpdateShopItem",
			}).Error(domain.ErrQuantityItems.Error())
			return domain.ShopItem{}, domain.ErrQuantityItems
		}
		shopItem.Quantity = param.Quantity.Int64
	}

	shopItem, err = s.shopRepo.UpdateShopItem(ctx, shopItem)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateShopItem",
		}).Error(err.Error())
		return domain.ShopItem{}, err
	}

	log.WithFields(log.Fields{
		"shopItemID": shopItem.ID,
		"Quantity":   shopItem.Quantity,
	}).Info("UpdateShopItem OK")
	return shopItem, nil
}

func (s *ShopService) DeleteShopItem(ctx context.Context, shopItemID domain.ID) error {
	err := s.shopRepo.DeleteShopItem(ctx, shopItemID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "DeleteShopItem",
		}).Error(err.Error())
		return err
	}

	log.WithFields(log.Fields{
		"shopItemID": shopItemID,
	}).Info("DeleteShopItem OK")
	return nil
}
