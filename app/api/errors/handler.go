package errors

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"

	libErr "github.com/futuramacoder/uniswap-api/app/pkg/errors"
)

func HandleError(c *gin.Context, err error) {
	var apiErr Error
	if errors.As(err, &apiErr) {
		abortWithError(c, apiErr)
		return
	}

	var domainErr libErr.Error
	if errors.As(err, &domainErr) {
		handleDomainError(c, domainErr)
		return
	}

	abortWithError(c, InternalError)
}

func handleDomainError(c *gin.Context, domainErr libErr.Error) {
	var err Error

	switch {
	case libErr.ErrorBadRequest.Is(domainErr):
		err = BadRequest.SetMessage(domainErr.Error())
	default:
		err = InternalError
	}

	abortWithError(c, err)
}

func abortWithError(c *gin.Context, err Error) {
	err.Timestamp = time.Now()
	err.Path = c.Request.URL.Path
	c.AbortWithStatusJSON(err.Status, &err)
}
