package postgres

import (
	"context"
	repository "github.com/EmirShimshir/marketplace/internal/adapter/repository/postgres"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var orderShopItems = []domain.OrderShopItem{
	domain.OrderShopItem{
		ID:          domain.ID("30e18bc1-4354-4937-9a3b-03cf0b70eee1"),
		OrderShopID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b702ee1"),
		ProductID:   domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027a1"),
		Quantity:    1,
	},
}

var orderShops = []domain.OrderShop{
	domain.OrderShop{
		ID:              domain.ID("30e18bc1-4354-4937-9a3b-03cf0b702ee1"),
		ShopID:          domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027b1"),
		OrderCustomerID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b702ae1"),
		Status:          domain.OrderShopStatusStart,
		Notified:        false,
		OrderShopItems:  orderShopItems,
	},
}

var orderCustomers = []domain.OrderCustomer{
	domain.OrderCustomer{
		ID:         domain.ID("30e18bc1-4354-4937-9a3b-03cf0b702ae1"),
		CustomerID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cc"),
		Address:    "Pushkina 1-2-3",
		CreatedAt:  time.Date(2022, 10, 10, 11, 30, 30, 0, time.UTC),
		OrderShops: orderShops,
	},
}

var createdOrderShopItems = []domain.OrderShopItem{
	domain.OrderShopItem{
		ID:          domain.ID("30e18bc1-4354-4937-9a3b-03cf0beeeeee"),
		OrderShopID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7eeeee"),
		ProductID:   domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027a1"),
		Quantity:    3,
	},
}

var createdorderShops = []domain.OrderShop{
	domain.OrderShop{
		ID:              domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7eeeee"),
		ShopID:          domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027b1"),
		OrderCustomerID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b70eeee"),
		Status:          domain.OrderShopStatusStart,
		Notified:        true,
		OrderShopItems:  createdOrderShopItems,
	},
}

var createdOrderCustomers = []domain.OrderCustomer{
	domain.OrderCustomer{
		ID:         domain.ID("30e18bc1-4354-4937-9a3b-03cf0b70eeee"),
		CustomerID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027cc"),
		Address:    "Pushkina 1-2-4",
		CreatedAt:  time.Date(2024, 10, 10, 11, 30, 30, 0, time.UTC),
		OrderShops: createdorderShops,
	},
}

func TestOrderRepository(t *testing.T) {
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

	t.Run("test GetOrderCustomerByID", func(t *testing.T) {
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

		repo := repository.NewOrderRepo(db)
		found, err := repo.GetOrderCustomerByID(ctx, orderCustomers[0].ID)
		if err != nil {
			t.Errorf("failed to GetOrderCustomerByID: %v", err)
		}

		require.Equal(t, orderCustomers[0], found)
	})
	t.Run("test CreateOrderCustomer", func(t *testing.T) {
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

		repo := repository.NewOrderRepo(db)
		found, err := repo.CreateOrderCustomer(ctx, createdOrderCustomers[0])
		if err != nil {
			t.Errorf("failed to CreateOrderCustomer: %v", err)
		}

		require.Equal(t, createdOrderCustomers[0], found)
	})
}
