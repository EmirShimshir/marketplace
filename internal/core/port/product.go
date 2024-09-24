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
	GetByID(ctx context.Context, productID domain.ID) (domain.Product, error)
}
