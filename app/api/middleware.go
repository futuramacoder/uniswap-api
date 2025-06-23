package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	gincors "github.com/rs/cors/wrapper/gin"

	apiErr "github.com/futuramacoder/uniswap-api/app/api/errors"
)

func (s *Server) loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.WithContext(c).WithFields(log.Fields{
			"ip":     c.ClientIP(),
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		}).Debug("request")

		start := time.Now()
		c.Next()

		log.WithContext(c).WithFields(log.Fields{
			"ip":      c.ClientIP(),
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
			"code":    c.Writer.Status(),
			"status":  http.StatusText(c.Writer.Status()),
			"latency": time.Now().Sub(start).String(),
		}).Debug("response")
	}
}

func (s *Server) corsMiddleware() gin.HandlerFunc {
	corsOptions := cors.Options{
		AllowedHeaders:     []string{"*"},
		AllowCredentials:   true,
		OptionsPassthrough: false,
		AllowedMethods:     []string{"HEAD", "GET", "POST", "PUT", "DELETE", "PATCH"},
		MaxAge:             1000,
	}

	if len(s.cfg.CorsOrigins) > 0 {
		// we use logrus instead of libs log because we don't have context here.
		log.WithField("origins", s.cfg.CorsOrigins).Debug("enabling cors origins")
		corsOptions.AllowedOrigins = s.cfg.CorsOrigins
	}

	if len(s.cfg.CorsMethods) > 0 {
		// we use logrus instead of libs log because we don't have context here.
		log.WithField("methods", s.cfg.CorsMethods).Debug("enabling cors methods")
		corsOptions.AllowedMethods = s.cfg.CorsMethods
	}

	return gincors.New(corsOptions)
}

func (s *Server) recoveryMiddleware() gin.HandlerFunc {
	return gin.RecoveryWithWriter(os.Stderr, func(c *gin.Context, err interface{}) {
		apiErr.HandleError(c, apiErr.InternalError)

		// recovery from panic logging
		defer func() {
			_ = recover()
		}()
		log.WithContext(c).WithError(fmt.Errorf("panic: %v", err)).Panic("recovery from panic")
	})
}
