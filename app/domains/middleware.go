package domains

import (
	"github.com/gin-gonic/gin"
	sdkTypes "github.com/h4lim/go-sdk/app/types"
)

type MiddlewareDelivery interface {
	GroupRouter(c *gin.Context)
}

type MiddlewareUseCase interface {
	CheckHeader() *sdkTypes.GeneralResponse
}
