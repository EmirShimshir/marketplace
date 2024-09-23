package postgres

import (
	"context"
	repository "github.com/EmirShimshir/marketplace/internal/adapter/repository/postgres"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

var cartItems = []domain.CartItem{
	domain.CartItem{
		ID:        domain.ID("30e18bc1-4354-4937-9a3b-03cf0b702aa1"),
		CartID:    domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cc"),
		ProductID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027a1"),
		Quantity:  2,
	},
	domain.CartItem{
		ID:        domain.ID("30e18bc1-4354-4937-9a3b-03cf0b702aa2"),
		CartID:    domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cc"),
		ProductID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027a2"),
		Quantity:  1,
	},
}

var createdCartItem = domain.CartItem{
	ID:        domain.ID("30e18bc1-4354-4937-9a3b-03cf0b702aa3"),
	CartID:    domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cd"),
	ProductID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027a2"),
	Quantity:  1,
}

var updatedCartItem = domain.CartItem{
	ID:        domain.ID("30e18bc1-4354-4937-9a3b-03cf0b702aa3"),
	CartID:    domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cd"),
	ProductID: domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7027a2"),
	Quantity:  2,
}

var carts = []domain.Cart{
	domain.Cart{
		ID:    domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cc"),
		Price: 0,
		Items: []domain.CartItem{
			cartItems[0],
			cartItems[1],
		},
	},
	domain.Cart{
		ID:    domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cd"),
		Price: 0,
		Items: []domain.CartItem{},
	},
}

var updatedCart = domain.Cart{
	ID:    domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cc"),
	Price: 2990,
	Items: []domain.CartItem{
		cartItems[0],
		cartItems[1],
	},
}

var clearedCart = domain.Cart{
	ID:    domain.ID("30e18bc1-4354-4937-9a3b-03cf0b7034cc"),
	Price: 2990,
	Items: []domain.CartItem{},
}

func TestCartRepository(t *testing.T) {
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

	t.Run("test get product 0", func(t *testing.T) {
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

		repo := repository.NewCartRepo(db)
		found, err := repo.GetCartByID(ctx, carts[0].ID)
		if err != nil {
			t.Errorf("failed to get cart: %v", err)
		}

		require.Equal(t, carts[0], found)
		for i := range carts[0].Items {
			require.Equal(t, carts[0].Items[i], found.Items[i])
		}
	})

	t.Run("test get product 1", func(t *testing.T) {
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

		repo := repository.NewCartRepo(db)
		found, err := repo.GetCartByID(ctx, carts[1].ID)
		if err != nil {
			t.Errorf("failed to get cart: %v", err)
		}

		require.Equal(t, carts[1], found)
		for i := range carts[1].Items {
			require.Equal(t, carts[1].Items[i], found.Items[i])
		}
	})

	t.Run("test update cart", func(t *testing.T) {
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

		repo := repository.NewCartRepo(db)
		cart, err := repo.UpdateCart(ctx, updatedCart)
		if err != nil {
			t.Errorf("failed to create cart: %v", err)
		}
		require.Equal(t, cart, updatedCart)
	})

	t.Run("test clear cart", func(t *testing.T) {
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

		repo := repository.NewCartRepo(db)
		err = repo.ClearCart(ctx, carts[0].ID)
		if err != nil {
			t.Errorf("failed to delete cart: %v", err)
		}
		found, err := repo.GetCartByID(ctx, carts[0].ID)
		if err != nil {
			t.Errorf("failed to get cart: %v", err)
		}
		require.Equal(t, found, clearedCart)
	})
	t.Run("test CreateCartItem", func(t *testing.T) {
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

		repo := repository.NewCartRepo(db)
		found, err := repo.CreateCartItem(ctx, createdCartItem)
		if err != nil {
			t.Errorf("failed to CreateCartItem: %v", err)
		}

		require.Equal(t, createdCartItem, found)
	})
	t.Run("test UpdateCartItem", func(t *testing.T) {
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

		repo := repository.NewCartRepo(db)

		found, err := repo.UpdateCartItem(ctx, updatedCartItem)
		if err != nil {
			t.Errorf("failed to UpdateCartItem: %v", err)
		}

		require.Equal(t, updatedCartItem, found)
	})
	t.Run("test DeleteCartItem", func(t *testing.T) {
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

		repo := repository.NewCartRepo(db)
		err = repo.DeleteCartItem(ctx, cartItems[0].ID)
		if err != nil {
			t.Errorf("failed to DeleteCartItem: %v", err)
		}
	})
}
