package postgres

import (
	"context"
	repository "github.com/EmirShimshir/marketplace/internal/adapter/repository/postgres"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/guregu/null"
	"github.com/stretchr/testify/require"
	"testing"
)

var users = []domain.User{
	domain.User{
		ID:       domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cb"),
		Name:     "Timur",
		Surname:  "Musin",
		Phone:    null.StringFrom("+79992233555"),
		Email:    "hanoys@mail.ru",
		Password: "qwerty",
		CartID:   domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cc"),
		Role:     domain.UserCustomer,
	},
	domain.User{
		ID:       domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cc"),
		Name:     "Emir",
		Surname:  "Shimshir",
		Phone:    null.String{},
		Email:    "emir@gmail.com",
		Password: "12345",
		CartID:   domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cd"),
		Role:     domain.UserCustomer,
	},
}

var createdUser = domain.User{
	ID:       domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cd"),
	CartID:   domain.ID("30e18bc1-4354-4937-9a3b-03cf0b111111"),
	Name:     "createdName",
	Surname:  "createdSurname",
	Phone:    null.StringFrom("+77777777777"),
	Email:    "user@mail.com",
	Password: "password",
	Role:     domain.UserCustomer,
}

var updatedUser = domain.User{
	ID:       domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cc"),
	CartID:   domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cd"),
	Name:     "Maxim",
	Surname:  "Shpakovsliy",
	Phone:    null.String{},
	Email:    "paw1a@yandex.ru",
	Password: "12345678",
	Role:     domain.UserCustomer,
}

func TestUserRepository(t *testing.T) {
	ctx := context.Background()
	container, err := newPostgresContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	url, err := container.ConnectionString(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("test get users", func(t *testing.T) {
		t.Cleanup(func() {
			err = container.Restore(ctx)
			if err != nil {
				t.Fatal(err)
			}
		})

		db, err := newPostgresDB(url)
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewUserRepo(db)
		found, err := repo.Get(ctx, 2, 0)
		if err != nil {
			t.Errorf("failed to get users: %v", err)
		}
		require.Equal(t, len(users), len(found))
		for i := range users {
			require.Equal(t, users[i], found[i])
		}
	})

	t.Run("test find user by id", func(t *testing.T) {
		t.Cleanup(func() {
			err = container.Restore(ctx)
			if err != nil {
				t.Fatal(err)
			}
		})

		db, err := newPostgresDB(url)
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewUserRepo(db)
		user, err := repo.GetByID(ctx, users[0].ID)
		if err != nil {
			t.Errorf("failed to find user with id: %v", err)
		}
		require.Equal(t, user, users[0])
	})

	t.Run("test create user", func(t *testing.T) {
		t.Cleanup(func() {
			err = container.Restore(ctx)
			if err != nil {
				t.Fatal(err)
			}
		})

		db, err := newPostgresDB(url)
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewUserRepo(db)
		user, err := repo.Create(ctx, createdUser)
		if err != nil {
			t.Errorf("failed to create user: %v", err)
		}
		require.Equal(t, user, createdUser)
	})

	t.Run("test update user", func(t *testing.T) {
		t.Cleanup(func() {
			err = container.Restore(ctx)
			if err != nil {
				t.Fatal(err)
			}
		})

		db, err := newPostgresDB(url)
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewUserRepo(db)
		user, err := repo.Update(ctx, updatedUser)
		if err != nil {
			t.Errorf("failed to create user: %v", err)
		}
		require.Equal(t, user, updatedUser)
	})

	t.Run("test delete user", func(t *testing.T) {
		t.Cleanup(func() {
			err = container.Restore(ctx)
			if err != nil {
				t.Fatal(err)
			}
		})

		db, err := newPostgresDB(url)
		if err != nil {
			t.Fatal(err)
		}

		repo := repository.NewUserRepo(db)
		err = repo.Delete(ctx, users[0].ID)
		if err != nil {
			t.Errorf("failed to delete user: %v", err)
		}
	})
}
