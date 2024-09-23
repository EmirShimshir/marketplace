package server

import (
	"fmt"
	v1 "github.com/EmirShimshir/marketplace/internal/adapter/delivery/http/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"

	_ "github.com/EmirShimshir/marketplace/docs"
)

func NewGinRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(cors.Default())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}

type ServerParams struct {
	Cfg     *v1.Config
	Handler *v1.Handler
	Router  *gin.Engine
}

func NewServer(params ServerParams) *http.Server {
	server := &http.Server{
		Handler:      params.Router,
		Addr:         fmt.Sprintf(":%s", params.Cfg.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return server
}
