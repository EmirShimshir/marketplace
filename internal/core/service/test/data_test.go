package service_test

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/guregu/null"
	"testing"
	"time"
)

var dataProducts = []domain.Product{
	{
		ID:          domain.NewID(),
		Name:        "Iphone 15 Pro",
		Description: "new",
		Price:       129990,
		Category:    domain.ElectronicCategory,
		PhotoUrl:    "http://photo1.png",
	},
	{
		ID:          domain.NewID(),
		Name:        "Harry Potter",
		Description: "new",
		Price:       2990,
		Category:    domain.BooksCategory,
		PhotoUrl:    "http://photo2.png",
	},
}

var dataCartsID = []domain.ID{
	domain.NewID(),
	domain.NewID(),
	domain.NewID(),
}

var dataCartItems = []domain.CartItem{
	{
		ID:        domain.NewID(),
		CartID:    dataCartsID[0],
		ProductID: dataProducts[0].ID,
		Quantity:  1,
	},
	{
		ID:        domain.NewID(),
		CartID:    dataCartsID[0],
		ProductID: dataProducts[1].ID,
		Quantity:  1,
	},
	{
		ID:        domain.NewID(),
		CartID:    dataCartsID[1],
		ProductID: dataProducts[0].ID,
		Quantity:  1,
	},
	{
		ID:        domain.NewID(),
		CartID:    dataCartsID[1],
		ProductID: dataProducts[1].ID,
		Quantity:  1,
	},
}

var dataCarts = []domain.Cart{
	{
		ID:    dataCartsID[0],
		Price: 0,
		Items: []domain.CartItem{
			dataCartItems[0],
			dataCartItems[1],
		},
	},
	{
		ID:    dataCartsID[1],
		Price: 0,
		Items: []domain.CartItem{
			dataCartItems[2],
			dataCartItems[3],
		},
	},
	{
		ID:    dataCartsID[2],
		Price: 0,
		Items: []domain.CartItem{},
	},
}

var dataUsers = []domain.User{
	{
		ID:       domain.NewID(),
		CartID:   dataCartsID[0],
		Name:     "Username1",
		Surname:  "Surname1",
		Phone:    null.StringFrom("+79150337964"),
		Email:    "user1@gmail.com",
		Password: "123456",
		Role:     domain.UserCustomer,
	},
	{
		ID:       domain.NewID(),
		CartID:   dataCartsID[1],
		Name:     "Username2",
		Surname:  "Surname2",
		Phone:    null.StringFrom("+7267793120"),
		Email:    "user2@gmail.com",
		Password: "654321",
		Role:     domain.UserCustomer,
	},
	{
		ID:       domain.NewID(),
		CartID:   dataCartsID[2],
		Name:     "Username3",
		Surname:  "Surname3",
		Phone:    null.StringFrom("+7267793121"),
		Email:    "user3@gmail.com",
		Password: "7654321",
		Role:     domain.UserCustomer,
	},
}

var dataShopsID = []domain.ID{
	domain.NewID(),
	domain.NewID(),
}

var dataShops = []domain.Shop{
	{
		ID:          dataShopsID[0],
		SellerID:    dataUsers[0].ID,
		Name:        "apple store",
		Description: "Description1",
		Requisites:  "Requisites1",
		Email:       "shop1@gmail.com",
		Items: []domain.ShopItem{
			dataShopItems[0],
		},
	},
	{
		ID:          dataShopsID[1],
		SellerID:    dataUsers[1].ID,
		Name:        "mybook",
		Description: "Description2",
		Requisites:  "Requisites2",
		Email:       "shop2@gmail.com",
		Items: []domain.ShopItem{
			dataShopItems[1],
		},
	},
}

var dataShopItems = []domain.ShopItem{
	{
		ID:        domain.NewID(),
		ShopID:    dataShopsID[0],
		ProductID: dataProducts[0].ID,
		Quantity:  4,
	},
	{
		ID:        domain.NewID(),
		ShopID:    dataShopsID[1],
		ProductID: dataProducts[1].ID,
		Quantity:  3,
	},
}

var dataWithdraws = []domain.Withdraw{
	{
		ID:      domain.NewID(),
		ShopID:  domain.NewID(),
		Comment: "comment",
		Sum:     1990,
		Status:  domain.WithdrawStatusStart,
	},
}

var dataOrderShopID = []domain.ID{
	domain.NewID(),
	domain.NewID(),
}

var dataOrderShopItems = []domain.OrderShopItem{
	{
		ID:          domain.NewID(),
		OrderShopID: dataOrderShopID[0],
		ProductID:   dataProducts[0].ID,
		Quantity:    1,
	},
	{
		ID:          domain.NewID(),
		OrderShopID: dataOrderShopID[1],
		ProductID:   dataProducts[1].ID,
		Quantity:    1,
	},
}

var dataOrderShops = []domain.OrderShop{
	{
		ID:     dataOrderShopID[0],
		ShopID: dataShopsID[0],
		Status: domain.OrderShopStatusStart,
		OrderShopItems: []domain.OrderShopItem{
			dataOrderShopItems[0],
		},
	},
	{
		ID:     dataOrderShopID[1],
		ShopID: dataShopsID[1],
		Status: domain.OrderShopStatusStart,
		OrderShopItems: []domain.OrderShopItem{
			dataOrderShopItems[1],
		},
	},
}

var dataOrderCustomers = []domain.OrderCustomer{
	{
		ID:         domain.NewID(),
		CustomerID: dataUsers[0].ID,
		Address:    "Pushkina 1-2-3",
		CreatedAt:  time.Now(),
		OrderShops: dataOrderShops,
	},
}

func TestData(t *testing.T) {}
