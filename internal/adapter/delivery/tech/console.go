package tech

import (
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	log "github.com/sirupsen/logrus"
	"os"
)

type Console struct {
	UserID  *domain.ID
	Handler *Handler
	Routes  map[Option]func(*Console)
}

type Option int

func NewConsole(handler *Handler) *Console {
	c := &Console{
		Handler: handler,
	}
	c.InitRoutes()
	return c
}

func (c *Console) InitRoutes() {
	c.Routes = map[Option]func(*Console){
		exit: func(c *Console) { os.Exit(0) },
		menu: func(c *Console) { c.PrintMenu() },

		signIn: c.Handler.SignIn,
		signUp: c.Handler.SignUp,
		logout: c.Handler.Logout,

		getUserAccount:    c.Handler.GetUser,
		updateUserAccount: c.Handler.UpdateUser,

		getShopItem:          c.Handler.GetShopItem,
		getShopItemsAll:      c.Handler.GetShopItemsAll,
		getShopItemsByShopID: c.Handler.GetShopItemsByShopID,
		getShopsBySellerID:   c.Handler.GetShopsBySellerID,
		createShop:           c.Handler.CreateShop,
		createShopItem:       c.Handler.CreateShopItem,
		updateShopItem:       c.Handler.UpdateShopItem,
		deleteShopItem:       c.Handler.DeleteShopItem,

		createCartItem: c.Handler.CreateCartItem,
		updateCartItem: c.Handler.UpdateCartItem,
		deleteCartItem: c.Handler.DeleteCartItem,
		getCart:        c.Handler.GetCart,

		createOrderCustomer:           c.Handler.CreateOrderCustomer,
		getOrderCustomersByCustomerID: c.Handler.GetOrderCustomersByCustomerID,
		getOrderShopsByShopID:         c.Handler.GetOrderShopsByShopID,
		updateOrderShopStatusByShopID: c.Handler.UpdateOrderShopStatusByShopID,

		getWithdrawsAll:     c.Handler.GetWithdrawsAll,
		getWithdrawByShopID: c.Handler.GetWithdrawByShopID,
		createWithdraw:      c.Handler.CreateWithdraw,
		updateWithdraw:      c.Handler.UpdateWithdraw,

		payOrder: c.Handler.PayOrder,
	}
}

const (
	exit = iota
	menu

	signIn
	signUp
	logout

	getUserAccount
	updateUserAccount

	getShopItem
	getShopItemsAll
	getShopItemsByShopID
	getShopsBySellerID
	createShop
	createShopItem
	updateShopItem
	deleteShopItem

	createCartItem
	updateCartItem
	deleteCartItem
	getCart

	createOrderCustomer
	getOrderCustomersByCustomerID
	getOrderShopsByShopID
	updateOrderShopStatusByShopID

	getWithdrawsAll
	getWithdrawByShopID
	createWithdraw
	updateWithdraw

	payOrder
)

func (c *Console) Start() {
	c.PrintMenu()
	for {
		var option Option
		fmt.Print("Введите команду: ")
		_, err := fmt.Scanf("%d", &option)
		if err != nil {
			fmt.Println("Ошибка ввода команды")
			continue
		}
		fmt.Println()

		handleFunc, ok := c.Routes[option]
		if !ok {
			fmt.Println("Ошибка ввода команды")
			continue
		}
		log.WithFields(log.Fields{
			"userID":  c.UserID,
			"command": option,
		}).Info("new command")
		handleFunc(c)
	}
}

func (c *Console) PrintMenu() {
	fmt.Println()
	fmt.Println("--------------------------------")
	fmt.Println("Меню (A - все; C - покупатель; S - продавец; M - модератор)")
	fmt.Println("0  Закрыть приложение")

	fmt.Println("1  Вывести меню")
	fmt.Println("2  Войти (G)")
	fmt.Println("3  Зарегистрироваться (G)")
	fmt.Println("4  Выйти (M | C | S)")

	fmt.Println("5  Получить информацию о пользователе (M | C | S)")
	fmt.Println("6  Обновить информацию о пользователе (M | C | S)")

	fmt.Println("7  Получить товар (G | M | C | S)")
	fmt.Println("8  Получить список всех товаров (G | M | C | S)")
	fmt.Println("9  Получить список товаров в магазине (G | M | C | S)")
	fmt.Println("10 Получить список магазинов продавца (S)")
	fmt.Println("11 Создать магазин (S)")
	fmt.Println("12 Добавить товар в магазин (S)")
	fmt.Println("13 Обновить товар в магазине (S)")
	fmt.Println("14 Удалить товар из магазина (M)")

	fmt.Println("15 Добавить товар в корзину (C)")
	fmt.Println("16 Изменить количество товара в корзине (C)")
	fmt.Println("17 Удалить товар из корзины (C)")
	fmt.Println("18 Получить список товаров в корзине (C)")

	fmt.Println("19 Создать заказ покупателя (С)")
	fmt.Println("20 Получить список заказов покупателя (С)")
	fmt.Println("21 Получить список заказов магазина (S)")
	fmt.Println("22 Изменить статус заказа магазина (S)")

	fmt.Println("23 Получить заявки на вывод средств (М)")
	fmt.Println("24 Получить заявки на вывод средств для магазина (S)")
	fmt.Println("25 Создать заявку на вывод средств (S)")
	fmt.Println("26 Обработать заявку на вывод средств (M)")

	fmt.Println("27 Оплатить заказ (C)")

	fmt.Println("--------------------------------")
}
