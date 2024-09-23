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

type PostgresProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *PostgresProductRepo {
	return &PostgresProductRepo{
		db: db,
	}
}

const (
	productGetQuery       = "SELECT * FROM public.product LIMIT $1 OFFSET $2"
	productGetByIDQuery   = "SELECT * FROM public.product WHERE id = $1"
	productGetByNameQuery = "SELECT * FROM public.product WHERE name = $1"
	productDeleteQuery    = "DELETE FROM public.product WHERE id = $1"
)

func (p *PostgresProductRepo) Get(ctx context.Context, limit, offset int64) ([]domain.Product, error) {
	var pgProducts []entity.PgProduct
	if err := p.db.SelectContext(ctx, &pgProducts, productGetQuery, limit, offset); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	products := make([]domain.Product, len(pgProducts))
	for i, product := range pgProducts {
		products[i] = product.ToDomain()
	}
	return products, nil
}

func (p *PostgresProductRepo) GetByID(ctx context.Context, productID domain.ID) (domain.Product, error) {
	var pgProduct entity.PgProduct
	if err := p.db.GetContext(ctx, &pgProduct, productGetByIDQuery, productID); err != nil {
		if err == sql.ErrNoRows {
			return domain.Product{}, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return domain.Product{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	return pgProduct.ToDomain(), nil
}

func (p *PostgresProductRepo) GetByName(ctx context.Context, name string) ([]domain.Product, error) {
	var pgProducts []entity.PgProduct
	if err := p.db.SelectContext(ctx, &pgProducts, productGetByNameQuery, name); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	products := make([]domain.Product, len(pgProducts))
	for i, product := range pgProducts {
		products[i] = product.ToDomain()
	}
	return products, nil
}

func (p *PostgresProductRepo) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	var pgProduct = entity.NewPgProduct(product)
	queryString := entity.InsertQueryString(pgProduct, "product")
	_, err := p.db.NamedExecContext(ctx, queryString, pgProduct)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.Product{}, errors.Wrap(domain.ErrDuplicate, err.Error())
			} else {
				return domain.Product{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.Product{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	return p.GetByID(ctx, product.ID)
}

func (p *PostgresProductRepo) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	var pgProduct = entity.NewPgProduct(product)
	queryString := entity.UpdateQueryString(pgProduct, "product")
	_, err := p.db.NamedExecContext(ctx, queryString, pgProduct)
	if err != nil {
		return domain.Product{}, errors.Wrap(domain.ErrUpdateFailed, err.Error())
	}

	return p.GetByID(ctx, product.ID)
}

func (p *PostgresProductRepo) Delete(ctx context.Context, productID domain.ID) error {
	_, err := p.db.ExecContext(ctx, productDeleteQuery, productID)
	if err != nil {
		return errors.Wrap(domain.ErrDeleteFailed, err.Error())
	}
	return nil
}
