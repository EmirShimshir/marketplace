package v1

import (
	"fmt"
	"github.com/EmirShimshir/marketplace/internal/core/domain"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"math"
	"net/http"
	"strconv"
	"time"
)

type Config struct {
	Host string
	Port string
}

type Handler struct {
	config          *Config
	shopService     port.IShopService
	productService  port.IProductService
	authService     port.IAuthService
	userService     port.IUserService
	cartService     port.ICartService
	orderService    port.IOrderService
	paymentService  port.IPayment
	withdrawService port.IWithdrawService
}

type HandlerParams struct {
	Config          *Config
	ShopService     port.IShopService
	ProductService  port.IProductService
	AuthService     port.IAuthService
	UserService     port.IUserService
	CartService     port.ICartService
	OrderService    port.IOrderService
	PaymentService  port.IPayment
	WithdrawService port.IWithdrawService
}

func NewHandler(params HandlerParams, router *gin.Engine) *Handler {
	handler := &Handler{
		config:          params.Config,
		shopService:     params.ShopService,
		productService:  params.ProductService,
		authService:     params.AuthService,
		userService:     params.UserService,
		cartService:     params.CartService,
		orderService:    params.OrderService,
		paymentService:  params.PaymentService,
		withdrawService: params.WithdrawService,
	}

	api := router.Group("/api")
	v1 := api.Group("/v1")
	v1.Use(LoggerMiddleware())
	{
		handler.initAuthRoutes(v1)
		handler.initUsersRoutes(v1)
		handler.initProductsRoutes(v1)
		handler.initShopsRoutes(v1)
		handler.initCartsRoutes(v1)
		handler.initOrdersRoutes(v1)
		handler.initWithdrawsRoutes(v1)
		handler.initPaymentsRoutes(v1)
	}

	return handler
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				if err.Type == gin.ErrorTypePrivate {
					log.Error(err.Error())
				}
			}
		} else {
			msg := fmt.Sprintf("[%s %d] %s (%dms)", c.Request.Method, statusCode, path, latency)
			if statusCode >= http.StatusInternalServerError {
				log.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				log.Warn(msg)
			} else {
				log.Info(msg)
			}
		}
	}
}

func getIdFromPath(c *gin.Context, paramName string) (domain.ID, error) {
	idString := c.Param(paramName)
	if idString == "" {
		return "", PathIdParamIsEmptyError
	}

	if _, err := uuid.Parse(idString); err != nil {
		return "", PathIdParamIsInvalidUUID
	}
	return domain.ID(idString), nil
}

func getQueryParamInt64(c *gin.Context, paramName string) (int64, error) {
	value := c.Query(paramName)
	if value == "" {
		return 0, PathIdParamIsEmptyError
	}
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, errors.Wrap(PathIdParamIsEmptyError, err.Error())
	}
	return i, nil
}

func getQueryParamString(c *gin.Context, paramName string) (string, error) {
	value := c.Query(paramName)
	if value == "" {
		return "", PathIdParamIsEmptyError
	}

	return value, nil
}

func getIdFromRequestContext(context *gin.Context) (domain.ID, error) {
	id, ok := context.Get("userID")
	if !ok {
		return "", UnauthorizedError
	}
	return domain.ID(id.(string)), nil
}

func getRoleFromRequestContext(context *gin.Context) (domain.UserRole, error) {
	role, ok := context.Get("role")
	if !ok {
		return 0, UnauthorizedError
	}
	return domain.UserRole(role.(int)), nil
}
