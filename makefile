include .env
export

build:
	go mod download && go build -o ./.bin/appTech ./cmd/tech/main.go
	go mod download && go build -o ./.bin/appWeb ./cmd/web/main.go

up:
	docker compose down
	docker compose up -d

test:
	go test -race ./...

copy_csv:
	export PGPASSWORD=postgres && psql -h localhost -p 5432 -U postgres -d postgres -f ./csv/copy_csv.sql

gen_data:
	python ./gen/gen_data.py
clear:
	docker compose down --volumes

run_tech: build
	./.bin/appTech

run_locust:
	locust -f ./resurch/resurch.py


run_web: build
	./.bin/appWeb

swagger:
	swag init --parseDependency  --parseInternal --parseDepth 1  -g ./cmd/web/main.go


migrate:
	# if "error: file does not exist" was occurred,
    # it means that data is up to date
	docker compose up migrate

mocks:
	mockery --dir internal/core/port --name IUserRepository --output internal/adapter/repository/mocks \
		--filename user.go --structname UserRepository
	mockery --dir internal/core/port --name IProductRepository --output internal/adapter/repository/mocks \
		--filename product.go --structname ProductRepository
	mockery --dir internal/core/port --name ICartRepository --output internal/adapter/repository/mocks \
		--filename cart.go --structname CartRepository
	mockery --dir internal/core/port --name IShopRepository --output internal/adapter/repository/mocks \
		--filename shop.go --structname ShopRepository
	mockery --dir internal/core/port --name IOrderRepository --output internal/adapter/repository/mocks \
		--filename order.go --structname OrderRepository
	mockery --dir internal/core/port --name IWithdrawRepository --output internal/adapter/repository/mocks \
		--filename withdraw.go --structname WithdrawRepository
	mockery --dir internal/core/port --name IUserService --output internal/core/service/mocks \
    		--filename user.go --structname UserService
	mockery --dir internal/core/port --name IProductService --output internal/core/service/mocks \
			--filename product.go --structname ProductService
	mockery --dir internal/core/port --name ICartService --output internal/core/service/mocks \
			--filename cart.go --structname CartService
	mockery --dir internal/core/port --name IShopService --output internal/core/service/mocks \
			--filename shop.go --structname ShopService
	mockery --dir internal/core/port --name IOrderService --output internal/core/service/mocks \
			--filename order.go --structname OrderService
	mockery --dir internal/core/port --name IWithdrawService --output internal/core/service/mocks \
			--filename withdraw.go --structname WithdrawService
	mockery --dir internal/core/port --name IEmailProvider --output internal/adapter/emailProvider/mocks \
			--filename email.go --structname EmailProvider
	mockery --dir internal/core/port --name IObjectStorage --output internal/adapter/storage/mocks \
			--filename storage.go --structname ObjectStorage
