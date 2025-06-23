package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/futuramacoder/uniswap-api/app/service/uniswap"
	_ "github.com/futuramacoder/uniswap-api/docs"
)

type Server struct {
	cfg        Config
	gin        *gin.Engine
	uniswapSvc *uniswap.Service
}

func NewServer(
	config Config,
	uniswapSvc *uniswap.Service,
) (s *Server, err error) {
	s = &Server{
		gin:        gin.New(),
		cfg:        config,
		uniswapSvc: uniswapSvc,
	}

	s.gin.Use(
		s.loggerMiddleware(),
		s.corsMiddleware(),
		s.recoveryMiddleware(),
	)

	s.configureHandlers()

	return s, nil
}

func (s *Server) Gin() *gin.Engine {
	return s.gin
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.cfg.Port)
	log.WithField("port", s.cfg.Port).Info("Starting api server")
	return s.gin.Run(addr)
}

// @title Uniswap API
// @version 0.0.1
// @description Uniswap API Documentation.
func (s *Server) configureHandlers() {
	generalGroup := s.gin.Group("/")
	generalGroup.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	generalGroup.GET("/estimate", s.estimateHandler())
}
