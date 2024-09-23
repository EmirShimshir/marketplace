package main

import (
	"github.com/EmirShimshir/marketplace/internal/app"
)

// @title        Marketplace API
// @version      1.0
// @description  This is simple api for marketplace

// @contact.name   API Support
// @contact.url    https://t.me/Emir_Shimshir
// @contact.email  emir2701@yandex.ru
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.RunWeb()
}
