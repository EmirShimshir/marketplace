package postgres

import (
	"context"
	repository "github.com/EmirShimshir/marketplace/internal/adapter/repository/postgres"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

var shopItems = []domain.ShopItem{
	domain.ShopItem{
		ID:        domain.ID("30e18bc1-4354-4937-9a3b-03cf0b702ac1"),
		ShopID:    domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027b1"),
		ProductID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027a1"),
		Quantity:  5,
	},
}

var shops = []domain.Shop{
	domain.Shop{
		ID:          domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027b1"),
		SellerID:    domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cb"),
		Name:        "Apple Store",
		Description: "found 1998",
		Requisites:  "Alabama",
		Email:       "Apple@mail.ru",
		Items:       shopItems,
	},
}

func TestShopRepository(t *testing.T) {
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

	t.Run("test GetShops", func(t *testing.T) {
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

		repo := repository.NewShopRepo(db)
		found, err := repo.GetShops(ctx, 2, 0)
		if err != nil {
			t.Errorf("failed to GetShops: %v", err)
		}

		require.Equal(t, len(shops), len(found))
		for i := range shops {
			require.Equal(t, shops[0], found[i])
		}
	})

	t.Run("test GetShopByID", func(t *testing.T) {
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

		repo := repository.NewShopRepo(db)
		found, err := repo.GetShopByID(ctx, shops[0].ID)
		if err != nil {
			t.Errorf("failed to GetShopByID: %v", err)
		}

		require.Equal(t, shops[0], found)
	})

	t.Run("test GetShopBySellerID", func(t *testing.T) {
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

		repo := repository.NewShopRepo(db)
		found, err := repo.GetShopBySellerID(ctx, shops[0].SellerID)
		if err != nil {
			t.Errorf("failed to GetShopByID: %v", err)
		}

		require.Equal(t, []domain.Shop{shops[0]}, found)
	})
}
