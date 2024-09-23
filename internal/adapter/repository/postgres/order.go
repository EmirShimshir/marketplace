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

type PostgresOrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *PostgresOrderRepo {
	return &PostgresOrderRepo{
		db: db,
	}
}

const (
	orderGetShopItemByProductIDQuery    = "SELECT * FROM public.shop_product WHERE product_id = $1"
	orderGetOrderShopByID               = "SELECT * FROM public.order_shop WHERE id = $1"
	orderGetOrderShopItemsByOrderShopID = "SELECT * FROM public.order_shop_product WHERE order_shop_id = $1"
	orderGetOrderCustomerByID           = "SELECT * FROM public.order_customer WHERE id = $1"
	orderGetOrderShopByOrderCustomerID  = "SELECT * FROM public.order_shop WHERE order_customer_id = $1"
	orderGetOrderCustomerByCustomerID   = "SELECT * FROM public.order_customer WHERE customer_id = $1"
	orderGetNoNotifiedOrderShops        = "SELECT * FROM public.order_shop WHERE notified = 'false'"
	orderGetOrderShopByShopID           = "SELECT * FROM public.order_shop WHERE shop_id = $1"
	orderUpdatePaymentStatus            = "UPDATE public.order_customer SET payed = 'true' WHERE id = $1"
)

func (o *PostgresOrderRepo) GetOrderCustomerByCustomerID(ctx context.Context, customerID domain.ID) ([]domain.OrderCustomer, error) {
	var pgOrderCustomers []entity.PgOrderCustomer
	if err := o.db.SelectContext(ctx, &pgOrderCustomers, orderGetOrderCustomerByCustomerID, customerID); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	orderCustomers := make([]domain.OrderCustomer, len(pgOrderCustomers))
	for i := range orderCustomers {
		orderCustomers[i] = pgOrderCustomers[i].ToDomain()
		var err error
		orderCustomers[i], err = o.GetOrderCustomerByID(ctx, orderCustomers[i].ID)
		if err != nil {
			return nil, err
		}
	}
	return orderCustomers, nil
}

func (o *PostgresOrderRepo) GetOrderCustomerByID(ctx context.Context, OrderCustomerID domain.ID) (domain.OrderCustomer, error) {
	var pgOrderCustomer entity.PgOrderCustomer
	if err := o.db.GetContext(ctx, &pgOrderCustomer, orderGetOrderCustomerByID, OrderCustomerID); err != nil {
		if err == sql.ErrNoRows {
			return domain.OrderCustomer{}, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return domain.OrderCustomer{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	orderCustomer := pgOrderCustomer.ToDomain()

	var pgOrderShops []entity.PgOrderShop
	if err := o.db.SelectContext(ctx, &pgOrderShops, orderGetOrderShopByOrderCustomerID, OrderCustomerID); err != nil {
		if err != sql.ErrNoRows {
			return domain.OrderCustomer{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	orderShops := make([]domain.OrderShop, len(pgOrderShops))
	for i := range orderShops {
		orderShops[i] = pgOrderShops[i].ToDomain()
		var err error
		orderShops[i].OrderShopItems, err = o.getOrderShopItemsByOrderShopID(ctx, orderShops[i].ID)
		if err != nil {
			return domain.OrderCustomer{}, err
		}
	}

	orderCustomer.OrderShops = orderShops

	return orderCustomer, nil
}

func (o *PostgresOrderRepo) getOrderShopItemsByOrderShopID(ctx context.Context, orderShopID domain.ID) ([]domain.OrderShopItem, error) {
	var pgOrderShopItems []entity.PgOrderShopItem
	if err := o.db.SelectContext(ctx, &pgOrderShopItems, orderGetOrderShopItemsByOrderShopID, orderShopID); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	orderShopItems := make([]domain.OrderShopItem, len(pgOrderShopItems))
	for i, orderShopItem := range pgOrderShopItems {
		orderShopItems[i] = orderShopItem.ToDomain()
	}

	return orderShopItems, nil
}
func (o *PostgresOrderRepo) GetOrderShopByID(ctx context.Context, orderShopID domain.ID) (domain.OrderShop, error) {
	var pgOrderShop entity.PgOrderShop
	if err := o.db.GetContext(ctx, &pgOrderShop, orderGetOrderShopByID, orderShopID); err != nil {
		if err == sql.ErrNoRows {
			return domain.OrderShop{}, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return domain.OrderShop{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	orderShopItems, err := o.getOrderShopItemsByOrderShopID(ctx, orderShopID)
	if err != nil {
		return domain.OrderShop{}, err
	}

	orderShop := pgOrderShop.ToDomain()
	orderShop.OrderShopItems = orderShopItems

	return orderShop, nil
}
func (o *PostgresOrderRepo) GetNoNotifiedOrderShops(ctx context.Context) ([]domain.OrderShop, error) {
	var pgOrderShops []entity.PgOrderShop
	if err := o.db.SelectContext(ctx, &pgOrderShops, orderGetNoNotifiedOrderShops); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	orderShops := make([]domain.OrderShop, len(pgOrderShops))
	for i := range orderShops {
		orderShops[i] = pgOrderShops[i].ToDomain()
		var err error
		orderShops[i].OrderShopItems, err = o.getOrderShopItemsByOrderShopID(ctx, orderShops[i].ID)
		if err != nil {
			return nil, err
		}
	}

	return orderShops, nil
}

func (o *PostgresOrderRepo) getPgEntities(orderCustomer domain.OrderCustomer) (entity.PgOrderCustomer, []entity.PgOrderShop, []entity.PgOrderShopItem) {
	pgOrderCustomer := entity.NewPgOrderCustomer(orderCustomer)
	pgOrderShops := make([]entity.PgOrderShop, 0)
	pgOrderShopItems := make([]entity.PgOrderShopItem, 0)
	for _, orderShop := range orderCustomer.OrderShops {
		pgOrderShops = append(pgOrderShops, entity.NewPgOrderShop(orderShop))
		for _, item := range orderShop.OrderShopItems {
			pgOrderShopItems = append(pgOrderShopItems, entity.NewPgOrderShopItem(item))
		}
	}
	return pgOrderCustomer, pgOrderShops, pgOrderShopItems
}

func (o *PostgresOrderRepo) txInsertOrderCustomer(ctx context.Context, tx *sqlx.Tx, pgOrderCustomer entity.PgOrderCustomer) error {
	queryString := entity.InsertQueryString(pgOrderCustomer, "order_customer")
	_, err := tx.NamedExecContext(ctx, queryString, pgOrderCustomer)
	if err != nil {
		tx.Rollback()
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return errors.Wrap(domain.ErrDuplicate, err.Error())
			} else {
				return errors.Wrap(domain.ErrPersistenceFailed, err.Error())
			}
		} else {
			return errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	return nil
}

func (o *PostgresOrderRepo) txInsertOrderShop(ctx context.Context, tx *sqlx.Tx, pgOrderShop entity.PgOrderShop) error {
	queryString := entity.InsertQueryString(pgOrderShop, "order_shop")
	_, err := tx.NamedExecContext(ctx, queryString, pgOrderShop)
	if err != nil {
		tx.Rollback()
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return errors.Wrap(domain.ErrDuplicate, err.Error())
			} else {
				return errors.Wrap(domain.ErrPersistenceFailed, err.Error())
			}
		} else {
			return errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	return nil
}

func (o *PostgresOrderRepo) txInsertOrderShopItem(ctx context.Context, tx *sqlx.Tx, pgOrderShopItem entity.PgOrderShopItem) error {
	queryString := entity.InsertQueryString(pgOrderShopItem, "order_shop_product")
	_, err := tx.NamedExecContext(ctx, queryString, pgOrderShopItem)
	if err != nil {
		tx.Rollback()
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return errors.Wrap(domain.ErrDuplicate, err.Error())
			} else {
				return errors.Wrap(domain.ErrPersistenceFailed, err.Error())
			}
		} else {
			return errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	return nil
}

func (o *PostgresOrderRepo) txUpdateShopItem(ctx context.Context, tx *sqlx.Tx, pgOrderShopItem entity.PgOrderShopItem) error {
	var pgShopItem entity.PgShopItem
	if err := tx.GetContext(ctx, &pgShopItem, orderGetShopItemByProductIDQuery, pgOrderShopItem.ProductID); err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			return errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	pgShopItem.Quantity -= pgOrderShopItem.Quantity
	queryString := entity.UpdateQueryString(pgShopItem, "shop_product")
	_, err := tx.NamedExecContext(ctx, queryString, pgShopItem)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(domain.ErrUpdateFailed, err.Error())
	}

	return nil
}

func (o *PostgresOrderRepo) CreateOrderCustomer(ctx context.Context, orderCustomer domain.OrderCustomer) (domain.OrderCustomer, error) {
	pgOrderCustomer, pgOrderShops, pgOrderShopItems := o.getPgEntities(orderCustomer)
	tx, err := o.db.Beginx()
	if err != nil {
		return domain.OrderCustomer{}, errors.Wrap(domain.ErrTransactionError, err.Error())
	}
	err = o.txInsertOrderCustomer(ctx, tx, pgOrderCustomer)
	if err != nil {
		return domain.OrderCustomer{}, err
	}
	for _, pgOrderShop := range pgOrderShops {
		err = o.txInsertOrderShop(ctx, tx, pgOrderShop)
		if err != nil {
			return domain.OrderCustomer{}, err
		}
	}
	for _, pgOrderShopItem := range pgOrderShopItems {
		err = o.txUpdateShopItem(ctx, tx, pgOrderShopItem)
		if err != nil {
			return domain.OrderCustomer{}, err
		}
		err = o.txInsertOrderShopItem(ctx, tx, pgOrderShopItem)
		if err != nil {
			return domain.OrderCustomer{}, err
		}
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return domain.OrderCustomer{}, errors.Wrap(domain.ErrTransactionError, err.Error())
	}
	return o.GetOrderCustomerByID(ctx, orderCustomer.ID)
}

func (o *PostgresOrderRepo) GetOrderShopByShopID(ctx context.Context, shopID domain.ID) ([]domain.OrderShop, error) {
	var pgOrderShops []entity.PgOrderShop
	if err := o.db.SelectContext(ctx, &pgOrderShops, orderGetOrderShopByShopID, shopID); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	orderShops := make([]domain.OrderShop, len(pgOrderShops))
	for i := range orderShops {
		orderShops[i] = pgOrderShops[i].ToDomain()
		var err error
		orderShops[i].OrderShopItems, err = o.getOrderShopItemsByOrderShopID(ctx, orderShops[i].ID)
		if err != nil {
			return nil, err
		}
	}

	return orderShops, nil
}
func (o *PostgresOrderRepo) UpdateOrderShop(ctx context.Context, orderShop domain.OrderShop) (domain.OrderShop, error) {
	var pgOrderShop = entity.NewPgOrderShop(orderShop)
	queryString := entity.UpdateQueryString(pgOrderShop, "order_shop")
	_, err := o.db.NamedExecContext(ctx, queryString, pgOrderShop)
	if err != nil {
		return domain.OrderShop{}, errors.Wrap(domain.ErrUpdateFailed, err.Error())
	}

	return o.GetOrderShopByID(ctx, orderShop.ID)
}

func (o *PostgresOrderRepo) UpdatePaymentStatus(ctx context.Context, orderCustomerID domain.ID) error {
	_, err := o.db.ExecContext(ctx, orderUpdatePaymentStatus, orderCustomerID)
	if err != nil {
		return errors.Wrap(domain.ErrUpdateFailed, err.Error())
	}
	return nil
}
