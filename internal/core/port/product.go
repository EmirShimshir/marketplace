package port

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/guregu/null"
	"io"
)

type CreateProductParam struct {
	Name        string
	Description string
	Price       int64
	Category    domain.ProductCategory
	PhotoReader io.Reader
}

type UpdateProductParam struct {
	Name        null.String
	Description null.String
	Price       null.Int
	Category    *domain.ProductCategory
	PhotoReader *io.Reader
}

type IProductService interface {
	Get(ctx context.Context, limit, offset int64) ([]domain.Product, error)
	GetByID(ctx context.Context, productID domain.ID) (domain.Product, error)
	GetByName(ctx context.Context, name string) ([]domain.Product, error)
	Create(ctx context.Context, param CreateProductParam) (domain.Product, error)
	Update(ctx context.Context, productID domain.ID, param UpdateProductParam) (domain.Product, error)
	Delete(ctx context.Context, productID domain.ID) error
}
