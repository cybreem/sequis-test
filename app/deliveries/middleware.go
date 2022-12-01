package deliveries

import (
	"strconv"
	"time"

	"github.com/cybreem/sequis-test/app"
	"github.com/cybreem/sequis-test/app/domains"
	"github.com/cybreem/sequis-test/app/usecases"
	"github.com/gin-gonic/gin"
	"github.com/h4lim/go-sdk/caching"
	"github.com/h4lim/go-sdk/utils"
)

type middlewareDeliveryContext struct {
}

func NewMiddlewareDelivery() domains.MiddlewareDelivery {
	return middlewareDeliveryContext{}
}

func (m middlewareDeliveryContext) GroupRouter(c *gin.Context) {
	now := time.Now().UTC()
	timestamp := now.UnixNano()
	app.SimpleCache.SetValue(app.CACHE_LOG, strconv.FormatInt(timestamp, 10), *caching.GlobalCache)
	middlewareUseCase := usecases.NewMiddlewareUseCase(c)
	responseError := middlewareUseCase.CheckHeader()

	if responseError != nil {
		gr := utils.GeneralResponseErrorBuild(time.Now().UTC(), responseError.Code,
			utils.GeneralMessageHandler("EN", responseError.Code))
		c.AbortWithStatusJSON(responseError.HttpCode, gin.H{
			"general_response": gr,
		})
	}
}
