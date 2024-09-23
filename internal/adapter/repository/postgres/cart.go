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

type PostgresCartRepo struct {
	db *sqlx.DB
}

func NewCartRepo(db *sqlx.DB) *PostgresCartRepo {
	return &PostgresCartRepo{
		db: db,
	}
}

const (
	cartGetByIDQuery             = "SELECT * FROM public.cart WHERE id = $1"
	cartGetCartItemsByIDQuery    = "SELECT * FROM public.cart_product WHERE cart_id = $1"
	cartItemsDeleteByCartIDQuery = "DELETE FROM public.cart_product WHERE cart_id = $1"
	cartItemGetByIQuery          = "SELECT * FROM public.cart_product WHERE id = $1"
	cartItemDeleteQuery          = "DELETE FROM public.cart_product WHERE id = $1"
)

func (c *PostgresCartRepo) GetCartByID(ctx context.Context, cartID domain.ID) (domain.Cart, error) {
	var pgCart entity.PgCart
	if err := c.db.GetContext(ctx, &pgCart, cartGetByIDQuery, cartID); err != nil {
		if err == sql.ErrNoRows {
			return domain.Cart{}, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return domain.Cart{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	var pgCartItems []entity.PgCartItem
	if err := c.db.SelectContext(ctx, &pgCartItems, cartGetCartItemsByIDQuery, cartID); err != nil {
		if err != sql.ErrNoRows {
			return domain.Cart{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	cartItems := make([]domain.CartItem, len(pgCartItems))
	for i, cartItem := range pgCartItems {
		cartItems[i] = cartItem.ToDomain()
	}

	cart := pgCart.ToDomain()
	cart.Items = cartItems

	return cart, nil
}

func (c *PostgresCartRepo) UpdateCart(ctx context.Context, cart domain.Cart) (domain.Cart, error) {
	var pgCart = entity.NewPgCart(cart)
	queryString := entity.UpdateQueryString(pgCart, "cart")
	_, err := c.db.NamedExecContext(ctx, queryString, pgCart)
	if err != nil {
		return domain.Cart{}, errors.Wrap(domain.ErrUpdateFailed, err.Error())
	}

	return c.GetCartByID(ctx, cart.ID)
}

func (c *PostgresCartRepo) ClearCart(ctx context.Context, cartID domain.ID) error {
	_, err := c.db.ExecContext(ctx, cartItemsDeleteByCartIDQuery, cartID)
	if err != nil {
		return errors.Wrap(domain.ErrDeleteFailed, err.Error())
	}
	return nil
}

func (c *PostgresCartRepo) GetCartItemByID(ctx context.Context, cartItemID domain.ID) (domain.CartItem, error) {
	var pgCartItem entity.PgCartItem
	if err := c.db.GetContext(ctx, &pgCartItem, cartItemGetByIQuery, cartItemID); err != nil {
		if err == sql.ErrNoRows {
			return domain.CartItem{}, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return domain.CartItem{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	return pgCartItem.ToDomain(), nil
}

func (c *PostgresCartRepo) CreateCartItem(ctx context.Context, cartItem domain.CartItem) (domain.CartItem, error) {
	var pgCartItem = entity.NewPgCartItem(cartItem)
	queryString := entity.InsertQueryString(pgCartItem, "cart_product")
	_, err := c.db.NamedExecContext(ctx, queryString, pgCartItem)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.CartItem{}, errors.Wrap(domain.ErrDuplicate, err.Error())
			} else {
				return domain.CartItem{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.CartItem{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	return c.GetCartItemByID(ctx, cartItem.ID)
}

func (c *PostgresCartRepo) UpdateCartItem(ctx context.Context, cartItem domain.CartItem) (domain.CartItem, error) {
	var pgCartItem = entity.NewPgCartItem(cartItem)
	queryString := entity.UpdateQueryString(pgCartItem, "cart_product")
	_, err := c.db.NamedExecContext(ctx, queryString, pgCartItem)
	if err != nil {
		return domain.CartItem{}, errors.Wrap(domain.ErrUpdateFailed, err.Error())
	}

	return c.GetCartItemByID(ctx, cartItem.ID)
}

func (c *PostgresCartRepo) DeleteCartItem(ctx context.Context, cartItemID domain.ID) error {
	_, err := c.db.ExecContext(ctx, cartItemDeleteQuery, cartItemID)
	if err != nil {
		return errors.Wrap(domain.ErrDeleteFailed, err.Error())
	}
	return nil
}
