package tech

import (
	"context"
	"fmt"
)

func (h *Handler) PayOrder(console *Console) {
	var key string
	fmt.Println("Введите ключ")
	fmt.Scanln(&key)

	err := h.payment.ProcessOrderPayment(context.Background(), key)
	if err != nil {
		ErrorResponse(err)
		return
	}

	fmt.Println("Заказ успешно оплачен")
}
