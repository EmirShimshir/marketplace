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

type PostgresUserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *PostgresUserRepo {
	return &PostgresUserRepo{
		db: db,
	}
}

const (
	userGetQuery        = "SELECT * FROM public.user LIMIT $1 OFFSET $2"
	userGetByIDQuery    = "SELECT * FROM public.user WHERE id = $1"
	userGetByEmailQuery = "SELECT * FROM public.user WHERE email = $1"
	userDeleteQuery     = "DELETE FROM public.user WHERE id = $1"
)

func (u *PostgresUserRepo) Get(ctx context.Context, limit, offset int64) ([]domain.User, error) {
	var pgUsers []entity.PgUser
	if err := u.db.SelectContext(ctx, &pgUsers, userGetQuery, limit, offset); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return nil, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	users := make([]domain.User, len(pgUsers))
	for i, user := range pgUsers {
		users[i] = user.ToDomain()
	}
	return users, nil
}

func (u *PostgresUserRepo) GetByID(ctx context.Context, userID domain.ID) (domain.User, error) {
	var pgUser entity.PgUser
	if err := u.db.GetContext(ctx, &pgUser, userGetByIDQuery, userID); err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return domain.User{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	return pgUser.ToDomain(), nil
}

func (u *PostgresUserRepo) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var pgUser entity.PgUser
	err := u.db.GetContext(ctx, &pgUser, userGetByEmailQuery, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.Wrap(domain.ErrNotExist, err.Error())
		} else {
			return domain.User{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	return pgUser.ToDomain(), nil
}

func (u *PostgresUserRepo) CreateWithoutCart(ctx context.Context, user domain.User) (domain.User, error) {
	var pgUser = entity.NewPgUser(user)
	queryString := entity.InsertQueryString(pgUser, "user")
	_, err := u.db.NamedExecContext(ctx, queryString, pgUser)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.User{}, errors.Wrap(domain.ErrDuplicate, err.Error())
			} else {
				return domain.User{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.User{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}

	return u.GetByID(ctx, user.ID)
}

func (u *PostgresUserRepo) Create(ctx context.Context, user domain.User) (domain.User, error) {
	tx, err := u.db.Beginx()
	if err != nil {
		return domain.User{}, errors.Wrap(domain.ErrTransactionError, err.Error())
	}
	var pgCart = entity.NewPgCart(domain.Cart{ID: user.CartID, Price: 0})
	queryString := entity.InsertQueryString(pgCart, "cart")
	_, err = tx.NamedExecContext(ctx, queryString, pgCart)
	if err != nil {
		tx.Rollback()
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.User{}, errors.Wrap(domain.ErrDuplicate, err.Error())
			} else {
				return domain.User{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.User{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	var pgUser = entity.NewPgUser(user)
	queryString = entity.InsertQueryString(pgUser, "user")
	_, err = tx.NamedExecContext(ctx, queryString, pgUser)
	if err != nil {
		tx.Rollback()
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == PgUniqueViolationCode {
				return domain.User{}, errors.Wrap(domain.ErrDuplicate, err.Error())
			} else {
				return domain.User{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
			}
		} else {
			return domain.User{}, errors.Wrap(domain.ErrPersistenceFailed, err.Error())
		}
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return domain.User{}, errors.Wrap(domain.ErrTransactionError, err.Error())
	}

	return u.GetByID(ctx, user.ID)
}

func (u *PostgresUserRepo) Update(ctx context.Context, user domain.User) (domain.User, error) {
	var pgUser = entity.NewPgUser(user)
	queryString := entity.UpdateQueryString(pgUser, "user")
	_, err := u.db.NamedExecContext(ctx, queryString, pgUser)
	if err != nil {
		return domain.User{}, errors.Wrap(domain.ErrUpdateFailed, err.Error())
	}

	return u.GetByID(ctx, user.ID)
}

func (u *PostgresUserRepo) Delete(ctx context.Context, userID domain.ID) error {
	_, err := u.db.ExecContext(ctx, userDeleteQuery, userID)
	if err != nil {
		return errors.Wrap(domain.ErrDeleteFailed, err.Error())
	}
	return nil
}
