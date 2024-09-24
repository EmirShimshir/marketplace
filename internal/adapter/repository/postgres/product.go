package postgres

import (
	"context"
	"database/sql"
	"github.com/EmirShimshir/marketplace/internal/adapter/repository/postgres/entity"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
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
	productGetByIDQuery = "SELECT * FROM public.product WHERE id = $1"
)

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
