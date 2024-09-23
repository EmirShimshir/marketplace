package app

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/adapter/authProvider/adapter/storage/redis"
	"github.com/EmirShimshir/marketplace/internal/adapter/authProvider/jwt"
	v1 "github.com/EmirShimshir/marketplace/internal/adapter/delivery/http/v1"
	"github.com/EmirShimshir/marketplace/internal/adapter/delivery/tech"
	emailProvider "github.com/EmirShimshir/marketplace/internal/adapter/emailProvider/goMailProvider"
	paymentMock "github.com/EmirShimshir/marketplace/internal/adapter/payment/mock"
	prostgresRepo "github.com/EmirShimshir/marketplace/internal/adapter/repository/postgres"
	storage "github.com/EmirShimshir/marketplace/internal/adapter/storage/minio"
	"github.com/EmirShimshir/marketplace/internal/app/config"
	"github.com/EmirShimshir/marketplace/internal/app/server"
	"github.com/EmirShimshir/marketplace/internal/core/service"
	"github.com/EmirShimshir/marketplace/pkg/database/postgres"
	redis2 "github.com/EmirShimshir/marketplace/pkg/database/redis"
	"github.com/EmirShimshir/marketplace/pkg/gomail"
	"github.com/EmirShimshir/marketplace/pkg/logrus"
	"github.com/EmirShimshir/marketplace/pkg/minio"
	log "github.com/sirupsen/logrus"
)

func RunTech() {
	cfg := config.GetConfig()
	logrus.InitLogrus(&cfg.Logger)
	log.Info("application startup...")

	db, err := postgres.NewPostgresDB(&cfg.Postgres)
	if err != nil {
		log.WithFields(log.Fields{
			"from":    "main()",
			"problem": "NewPostgresDB",
		}).Fatal(err.Error())
	}
	shopRepo := prostgresRepo.NewShopRepo(db)
	productRepo := prostgresRepo.NewProductRepo(db)
	userRepo := prostgresRepo.NewUserRepo(db)
	cartRepo := prostgresRepo.NewCartRepo(db)
	orderRepo := prostgresRepo.NewOrderRepo(db)
	withdrawRepo := prostgresRepo.NewWithdrawRepo(db)

	minioClient, err := minio.NewClient(&cfg.Minio)
	if err != nil {
		log.WithFields(log.Fields{
			"from":    "main()",
			"problem": "NewMinioClient",
		}).Fatal(err.Error())
	}
	objectStorage := storage.NewObjectStorage(minioClient, &cfg.Minio)
	shopService := service.NewShopService(shopRepo, objectStorage)

	productService := service.NewProductService(productRepo, objectStorage)

	userService := service.NewUserService(userRepo)

	redisClient, err := redis2.NewClient(&cfg.Redis)
	if err != nil {
		log.WithFields(log.Fields{
			"from":    "main()",
			"problem": "NewRedisClient",
		}).Fatal(err.Error())
	}
	sessionStorage := redis.NewSessionStorage(redisClient)
	authProvider := jwt.NewAuthProvider(&cfg.JWT, sessionStorage)
	authService := service.NewAuthService(authProvider, userService)

	cartService := service.NewCartService(cartRepo, shopRepo, productRepo)

	orderService := service.NewOrderService(orderRepo, userRepo, cartRepo, shopRepo)
	gateway := paymentMock.NewtMockGateway()
	paymentService := service.NewPaymentService(gateway, orderRepo)
	withdrawService := service.NewWithdrawService(withdrawRepo)
	params := tech.HandlerParams{
		Shop:     shopService,
		Product:  productService,
		Auth:     authService,
		User:     userService,
		Cart:     cartService,
		Order:    orderService,
		Payment:  paymentService,
		Withdraw: withdrawService,
	}
	h := tech.NewHandler(params)
	c := tech.NewConsole(h)

	goMailClient := gomail.NewClient(&cfg.Gomail)
	goMailProvider := emailProvider.NewGoMailProvider(goMailClient)
	emailService := service.NewEmailService(goMailProvider, orderRepo, shopRepo)
	stopCh := make(chan struct{})

	log.Info("emailNotifier started")
	go emailService.Start(context.Background(), stopCh)

	log.Info("app started")
	c.Start()
	stopCh <- struct{}{}
}

func RunWeb() {
	cfg := config.GetConfig()
	logrus.InitLogrus(&cfg.Logger)
	log.Info("application startup...")

	db, err := postgres.NewPostgresDB(&cfg.Postgres)
	if err != nil {
		log.WithFields(log.Fields{
			"from":    "main()",
			"problem": "NewPostgresDB",
		}).Fatal(err.Error())
	}
	shopRepo := prostgresRepo.NewShopRepo(db)
	productRepo := prostgresRepo.NewProductRepo(db)
	userRepo := prostgresRepo.NewUserRepo(db)
	cartRepo := prostgresRepo.NewCartRepo(db)
	orderRepo := prostgresRepo.NewOrderRepo(db)
	withdrawRepo := prostgresRepo.NewWithdrawRepo(db)

	minioClient, err := minio.NewClient(&cfg.Minio)
	if err != nil {
		log.WithFields(log.Fields{
			"from":    "main()",
			"problem": "NewMinioClient",
		}).Fatal(err.Error())
	}
	objectStorage := storage.NewObjectStorage(minioClient, &cfg.Minio)

	shopService := service.NewShopService(shopRepo, objectStorage)

	productService := service.NewProductService(productRepo, objectStorage)

	userService := service.NewUserService(userRepo)

	redisClient, err := redis2.NewClient(&cfg.Redis)
	if err != nil {
		log.WithFields(log.Fields{
			"from":    "main()",
			"problem": "NewRedisClient",
		}).Fatal(err.Error())
	}
	sessionStorage := redis.NewSessionStorage(redisClient)
	authProvider := jwt.NewAuthProvider(&cfg.JWT, sessionStorage)
	authService := service.NewAuthService(authProvider, userService)

	cartService := service.NewCartService(cartRepo, shopRepo, productRepo)

	orderService := service.NewOrderService(orderRepo, userRepo, cartRepo, shopRepo)
	gateway := paymentMock.NewtMockGateway()
	paymentService := service.NewPaymentService(gateway, orderRepo)
	withdrawService := service.NewWithdrawService(withdrawRepo)

	handlerParams := v1.HandlerParams{
		Config:          &cfg.Web,
		ShopService:     shopService,
		ProductService:  productService,
		AuthService:     authService,
		UserService:     userService,
		CartService:     cartService,
		OrderService:    orderService,
		PaymentService:  paymentService,
		WithdrawService: withdrawService,
	}
	gin := server.NewGinRouter()
	v1 := v1.NewHandler(handlerParams, gin)

	serverParams := server.ServerParams{
		Cfg:     &cfg.Web,
		Handler: v1,
		Router:  gin,
	}
	server := server.NewServer(serverParams)

	goMailClient := gomail.NewClient(&cfg.Gomail)
	goMailProvider := emailProvider.NewGoMailProvider(goMailClient)
	emailService := service.NewEmailService(goMailProvider, orderRepo, shopRepo)
	stopCh := make(chan struct{})

	log.Info("emailNotifier started")
	go emailService.Start(context.Background(), stopCh)

	log.Info("app started")
	log.Fatal("server shutdown", server.ListenAndServe())
	stopCh <- struct{}{}
}
