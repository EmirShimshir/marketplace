package service

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	log "github.com/sirupsen/logrus"
)

type ShopService struct {
	repo    port.IShopRepository
	storage port.IObjectStorage
}

func NewShopService(repo port.IShopRepository, storage port.IObjectStorage) *ShopService {
	return &ShopService{
		repo:    repo,
		storage: storage,
	}
}

func (s *ShopService) GetShops(ctx context.Context, limit, offset int64) ([]domain.Shop, error) {
	shop, err := s.repo.GetShops(ctx, limit, offset)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetShops",
		}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{
		"count": len(shop),
	}).Info("GetShops OK")
	return shop, nil
}

func (s *ShopService) GetShopByID(ctx context.Context, shopID domain.ID) (domain.Shop, error) {
	shop, err := s.repo.GetShopByID(ctx, shopID)
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
	shop, err := s.repo.GetShopBySellerID(ctx, sellerID)
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

	shop, err := s.repo.CreateShop(ctx, domain.Shop{
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

func (s *ShopService) UpdateShop(ctx context.Context, shopID domain.ID, param port.UpdateShopParam) (domain.Shop, error) {
	shop, err := s.GetShopByID(ctx, shopID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateShop",
		}).Error(err.Error())
		return domain.Shop{}, err
	}

	if param.Name.Valid {
		shop.Name = param.Name.String
	}
	if param.Description.Valid {
		shop.Description = param.Description.String
	}
	if param.Requisites.Valid {
		shop.Requisites = param.Requisites.String
	}
	if param.Email.Valid {
		shop.Email = param.Email.String
	}

	if shop.Name == "" {
		return domain.Shop{}, domain.ErrName
	}
	if shop.Description == "" {
		return domain.Shop{}, domain.ErrDescription
	}
	if shop.Requisites == "" {
		return domain.Shop{}, domain.ErrRequisites
	}

	shop, err = s.repo.UpdateShop(ctx, shop)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "UpdateShop",
		}).Error(err.Error())
		return domain.Shop{}, err
	}

	log.WithFields(log.Fields{
		"SellerID": shop.SellerID,
		"Email":    shop.Email,
	}).Info("UpdateShop OK")
	return shop, nil
}

func (s *ShopService) DeleteShop(ctx context.Context, shopID domain.ID) error {
	err := s.repo.DeleteShop(ctx, shopID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "DeleteShop",
		}).Error(err.Error())
		return err
	}

	log.WithFields(log.Fields{
		"shopID": shopID,
	}).Info("DeleteShop OK")
	return nil
}

func (s *ShopService) GetShopItemByID(ctx context.Context, shopItemID domain.ID) (domain.ShopItem, error) {
	shopItem, err := s.repo.GetShopItemByID(ctx, shopItemID)
	if err != nil {
		log.WithFields(log.Fields{
			"from": "GetShopItemByID",
		}).Error(err.Error())
		return domain.ShopItem{}, err
	}

	log.WithFields(log.Fields{
		"shopItemID": shopItemID,
	}).Info("GetShopItemByID OK")
	return shopItem, nil
}

func (s *ShopService) GetShopItems(ctx context.Context, limit, offset int64) ([]domain.ShopItem, error) {
	shopItems, err := s.repo.GetShopItems(ctx, limit, offset)
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
	shopItem, err := s.repo.GetShopItemByProductID(ctx, productID)
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

	shopItem, err = s.repo.CreateShopItem(ctx, shopItem, product)
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
	shopItem, err := s.repo.GetShopItemByID(ctx, shopItemID)
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

	shopItem, err = s.repo.UpdateShopItem(ctx, shopItem)
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
	err := s.repo.DeleteShopItem(ctx, shopItemID)
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
