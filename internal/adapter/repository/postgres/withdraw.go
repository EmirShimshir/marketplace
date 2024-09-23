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

type PostgresWithdrawRepo struct {
	db *sqlx.DB
}

func NewWithdrawRepo(db *sqlx.DB) *PostgresWithdrawRepo {
	return &PostgresWithdrawRepo{
		db: db,
	}
}

const (
	withdrawGetQuery         = "SELECT * FROM public.withdraw LIMIT $1 OFFSET $2"
	withdrawGetByIDQuery     = "SELECT * FROM public.withdraw WHERE id = $1"
	withdrawGetByShopIDQuery = "SELECT * FROM public.withdraw WHERE shop_id = $1"
	WithdrawDeleteQuery      = "SELECT * FROM public.withdraw WHERE id = $1"
)

func (w *PostgresWithdrawRepo) Get(ctx context.Context, limit, offset int64) ([]domain.Withdraw, error) {
	var pgWithdraws []entity.PgWithdraw
	if err := w.db.SelectContext(ctx, &pgWithdraws, withdrawGetQuery, limit, offset); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	withdraws := make([]domain.Withdraw, len(pgWithdraws))
	for i, withdraw := range pgWithdraws {
		withdraws[i] = withdraw.ToDomain()
	}
	return withdraws, nil
}

func (w *PostgresWithdrawRepo) GetByID(ctx context.Context, WithdrawID domain.ID) (domain.Withdraw, error) {
	var pgWithdraw entity.PgWithdraw
	if err := w.db.GetContext(ctx, &pgWithdraw, withdrawGetByIDQuery, WithdrawID); err != nil {
		if err == sql.ErrNoRows {
			return domain.Withdraw{}, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return domain.Withdraw{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	return pgWithdraw.ToDomain(), nil
}
func (w *PostgresWithdrawRepo) GetByShopID(ctx context.Context, shopID domain.ID) ([]domain.Withdraw, error) {
	var pgWithdraws []entity.PgWithdraw
	if err := w.db.SelectContext(ctx, &pgWithdraws, withdrawGetByShopIDQuery, shopID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	withdraws := make([]domain.Withdraw, len(pgWithdraws))
	for i, withdraw := range pgWithdraws {
		withdraws[i] = withdraw.ToDomain()
	}
	return withdraws, nil
}
func (w *PostgresWithdrawRepo) Create(ctx context.Context, withdraw domain.Withdraw) (domain.Withdraw, error) {
	var pgWithdraw = entity.NewPgWithdraw(withdraw)
	queryString := entity.InsertQueryString(pgWithdraw, "withdraw")
	_, err := w.db.NamedExecContext(ctx, queryString, pgWithdraw)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.Withdraw{}, errors.Wrap(domain.ErrDuplicate, err.Error())
			} else {
				return domain.Withdraw{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.Withdraw{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	return w.GetByID(ctx, withdraw.ID)
}

func (w *PostgresWithdrawRepo) Update(ctx context.Context, withdraw domain.Withdraw) (domain.Withdraw, error) {
	var pgWithdraw = entity.NewPgWithdraw(withdraw)
	queryString := entity.UpdateQueryString(pgWithdraw, "withdraw")
	_, err := w.db.NamedExecContext(ctx, queryString, pgWithdraw)
	if err != nil {
		return domain.Withdraw{}, errors.Wrap(domain.ErrUpdateFailed, err.Error())
	}

	return w.GetByID(ctx, withdraw.ID)
}
func (w *PostgresWithdrawRepo) Delete(ctx context.Context, withdrawID domain.ID) error {
	_, err := w.db.ExecContext(ctx, WithdrawDeleteQuery, withdrawID)
	if err != nil {
		return errors.Wrap(domain.ErrDeleteFailed, err.Error())
	}
	return nil
}
