package api

import (
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/futuramacoder/uniswap-api/app/api/errors"
	"github.com/futuramacoder/uniswap-api/app/api/payload"
)

// estimateHandler godoc
// @Summary Estimate Uniswap V2 pool
// @Description Estimate Uniswap V2 pool
// @Tags Estimate
// @Produce json
// @Param pool query string true "pool"
// @Param src query string true "src"
// @Param dst query string true "dst"
// @Param src_amount query string true "src_amount"
// @Success 200 {object} payload.EstimateResponse
// @Failure 400 {object} errors.Error
// @Failure 401 {object} errors.Error
// @Failure 404 {object} errors.Error
// @Failure 500 {object} errors.Error
// @Router /estimate [get]
func (s *Server) estimateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.WithContext(ctx).Debug("uniswap v2 estimate handler")
		pool := ctx.Query("pool")
		if pool == "" {
			errors.HandleError(ctx, errors.ValidationError.SetMessage("pool required field"))
			return
		}
		src := ctx.Query("src")
		if src == "" {
			errors.HandleError(ctx, errors.ValidationError.SetMessage("src required field"))
			return
		}
		dst := ctx.Query("dst")
		if dst == "" {
			errors.HandleError(ctx, errors.ValidationError.SetMessage("dst required field"))
			return
		}
		srcAmount := ctx.Query("src_amount")
		if srcAmount == "" {
			errors.HandleError(ctx, errors.ValidationError.SetMessage("src_amount required field"))
			return
		}

		amount, _ := new(big.Int).SetString(srcAmount, 0)
		outputAmount, err := s.uniswapSvc.GetOutputAmount(ctx, src, dst, pool, amount)
		if err != nil {
			errors.HandleError(ctx, errors.BadRequest.SetMessage(err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, payload.EstimateResponse{Amount: outputAmount.String()})
	}
}
