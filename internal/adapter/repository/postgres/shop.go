package postgres

import (
	"context"
	"database/sql"
	"github.com/EmirShimshir/marketplace/internal/adapter/repository/postgres/entity"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/jackc/pgconn"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type PostgresShopRepo struct {
	db *sqlx.DB
}

func NewShopRepo(db *sqlx.DB) *PostgresShopRepo {
	return &PostgresShopRepo{
		db: db,
	}
}

const (
	shopGetQuery                = "SELECT * FROM public.shop LIMIT $1 OFFSET $2"
	shopGetByIDQuery            = "SELECT * FROM public.shop WHERE id = $1"
	shopGetBySellerIDQuery      = "SELECT * FROM public.shop WHERE seller_id = $1"
	shopDeleteQuery             = "DELETE FROM public.shop WHERE id = $1"
	shopItemsGetQuery           = "SELECT * FROM public.shop_product LIMIT $1 OFFSET $2"
	shopItemGetByIDQuery        = "SELECT * FROM public.shop_product WHERE id = $1"
	shopItemGetByProductIDQuery = "SELECT * FROM public.shop_product WHERE product_id = $1"
	shopItemsGetByShopID        = "SELECT * FROM public.shop_product WHERE shop_id = $1"
	shopItemDeleteQuery         = "DELETE FROM public.shop_product WHERE id = $1"
)

func (o *PostgresShopRepo) GetShops(ctx context.Context, limit, offset int64) ([]domain.Shop, error) {
	var pgShops []entity.PgShop
	if err := o.db.SelectContext(ctx, &pgShops, shopGetQuery, limit, offset); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	shops := make([]domain.Shop, len(pgShops))
	for i, shop := range pgShops {
		shops[i] = shop.ToDomain()
		shopItems, err := o.getShopItemsByShopID(ctx, shops[i].ID)
		if err != nil {
			return nil, err
		}
		shops[i].Items = shopItems

	}
	return shops, nil
}

func (o *PostgresShopRepo) GetShopByID(ctx context.Context, shopID domain.ID) (domain.Shop, error) {
	var pgShop entity.PgShop
	if err := o.db.GetContext(ctx, &pgShop, shopGetByIDQuery, shopID); err != nil {
		if err == sql.ErrNoRows {
			return domain.Shop{}, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return domain.Shop{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	shop := pgShop.ToDomain()
	shopItems, err := o.getShopItemsByShopID(ctx, shop.ID)
	if err != nil {
		return domain.Shop{}, err
	}
	shop.Items = shopItems

	return shop, nil
}

func (o *PostgresShopRepo) GetShopBySellerID(ctx context.Context, sellerID domain.ID) ([]domain.Shop, error) {
	var pgShops []entity.PgShop
	if err := o.db.SelectContext(ctx, &pgShops, shopGetBySellerIDQuery, sellerID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	shops := make([]domain.Shop, len(pgShops))
	for i, shop := range pgShops {
		shops[i] = shop.ToDomain()
		shopItems, err := o.getShopItemsByShopID(ctx, shops[i].ID)
		if err != nil {
			return nil, err
		}
		shops[i].Items = shopItems

	}
	return shops, nil
}

func (o *PostgresShopRepo) CreateShop(ctx context.Context, shop domain.Shop) (domain.Shop, error) {
	var pgShop = entity.NewPgShop(shop)
	queryString := entity.InsertQueryString(pgShop, "shop")
	_, err := o.db.NamedExecContext(ctx, queryString, pgShop)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.Shop{}, errors.Wrap(domain.ErrDuplicate, err.Error())
			} else {
				return domain.Shop{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.Shop{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	return o.GetShopByID(ctx, shop.ID)
}

func (o *PostgresShopRepo) UpdateShop(ctx context.Context, shop domain.Shop) (domain.Shop, error) {
	var pgShop = entity.NewPgShop(shop)
	queryString := entity.UpdateQueryString(pgShop, "shop")
	_, err := o.db.NamedExecContext(ctx, queryString, pgShop)
	if err != nil {
		return domain.Shop{}, errors.Wrap(domain.ErrUpdateFailed, err.Error())
	}

	return o.GetShopByID(ctx, shop.ID)
}
func (o *PostgresShopRepo) DeleteShop(ctx context.Context, shopID domain.ID) error {
	_, err := o.db.ExecContext(ctx, shopDeleteQuery, shopID)
	if err != nil {
		return errors.Wrap(domain.ErrDeleteFailed, err.Error())
	}
	return nil
}

func (o *PostgresShopRepo) GetShopItems(ctx context.Context, limit, offset int64) ([]domain.ShopItem, error) {
	var pgShopItems []entity.PgShopItem
	if err := o.db.SelectContext(ctx, &pgShopItems, shopItemsGetQuery, limit, offset); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	shopItems := make([]domain.ShopItem, len(pgShopItems))
	for i, item := range pgShopItems {
		shopItems[i] = item.ToDomain()
	}
	return shopItems, nil
}
func (o *PostgresShopRepo) GetShopItemByID(ctx context.Context, shopItemID domain.ID) (domain.ShopItem, error) {
	var pgShopItem entity.PgShopItem
	if err := o.db.GetContext(ctx, &pgShopItem, shopItemGetByIDQuery, shopItemID); err != nil {
		if err == sql.ErrNoRows {
			return domain.ShopItem{}, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return domain.ShopItem{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	return pgShopItem.ToDomain(), nil
}
func (o *PostgresShopRepo) GetShopItemByProductID(ctx context.Context, productID domain.ID) (domain.ShopItem, error) {
	var pgShopItem entity.PgShopItem
	if err := o.db.GetContext(ctx, &pgShopItem, shopItemGetByProductIDQuery, productID); err != nil {
		if err == sql.ErrNoRows {
			return domain.ShopItem{}, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return domain.ShopItem{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	return pgShopItem.ToDomain(), nil
}

func (o *PostgresShopRepo) CreateShopItem(ctx context.Context, shopItem domain.ShopItem, product domain.Product) (domain.ShopItem, error) {
	tx, err := o.db.Beginx()
	if err != nil {
		return domain.ShopItem{}, errors.Wrap(domain.ErrTransactionError, err.Error())
	}

	var pgProduct = entity.NewPgProduct(product)
	queryString := entity.InsertQueryString(pgProduct, "product")
	_, err = tx.NamedExecContext(ctx, queryString, pgProduct)
	if err != nil {
		tx.Rollback()
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.ShopItem{}, errors.Wrap(domain.ErrDuplicate, err.Error())
			} else {
				return domain.ShopItem{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.ShopItem{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	var pgShopItem = entity.NewPgShopItem(shopItem)
	queryString = entity.InsertQueryString(pgShopItem, "shop_product")
	_, err = tx.NamedExecContext(ctx, queryString, pgShopItem)
	if err != nil {
		tx.Rollback()
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.ShopItem{}, errors.Wrap(domain.ErrDuplicate, err.Error())
			} else {
				return domain.ShopItem{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.ShopItem{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return domain.ShopItem{}, errors.Wrap(domain.ErrTransactionError, err.Error())
	}

	return o.GetShopItemByID(ctx, shopItem.ID)
}

func (o *PostgresShopRepo) UpdateShopItem(ctx context.Context, shopItem domain.ShopItem) (domain.ShopItem, error) {
	var pgShopItem = entity.NewPgShopItem(shopItem)
	queryString := entity.UpdateQueryString(pgShopItem, "shop_product")
	_, err := o.db.NamedExecContext(ctx, queryString, pgShopItem)
	if err != nil {
		return domain.ShopItem{}, errors.Wrap(domain.ErrUpdateFailed, err.Error())
	}

	return o.GetShopItemByID(ctx, shopItem.ID)
}

func (o *PostgresShopRepo) DeleteShopItem(ctx context.Context, shopItemID domain.ID) error {
	_, err := o.db.ExecContext(ctx, shopItemDeleteQuery, shopItemID)
	if err != nil {
		return errors.Wrap(domain.ErrDeleteFailed, err.Error())
	}
	return nil
}

func (o *PostgresShopRepo) getShopItemsByShopID(ctx context.Context, shopID domain.ID) ([]domain.ShopItem, error) {
	var pgShopItems []entity.PgShopItem
	if err := o.db.SelectContext(ctx, &pgShopItems, shopItemsGetByShopID, shopID); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	shopItems := make([]domain.ShopItem, len(pgShopItems))
	for i, item := range pgShopItems {
		shopItems[i] = item.ToDomain()
	}
	return shopItems, nil
}
