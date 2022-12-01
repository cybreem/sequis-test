package usecases

import (
	"fmt"
	"net/http"

	"github.com/cybreem/sequis-test/app"
	"github.com/cybreem/sequis-test/app/domains"
	"github.com/cybreem/sequis-test/app/domains/types"
	"github.com/gin-gonic/gin"
	sdkTypes "github.com/h4lim/go-sdk/app/types"
	"github.com/h4lim/go-sdk/caching"
	"github.com/h4lim/go-sdk/utils"
)

type middlewareUseCaseContext struct {
	ginContext *gin.Context
}

func NewMiddlewareUseCase(ginContext *gin.Context) domains.MiddlewareUseCase {
	return middlewareUseCaseContext{ginContext: ginContext}
}

func (m middlewareUseCaseContext) CheckHeader() *sdkTypes.GeneralResponse {

	var header types.HeaderContent
	if err := m.ginContext.ShouldBindHeader(&header); err != nil {
		// 	return utils.GetResponseMessage("EN",
		// 		http.StatusBadRequest, 99)
	}

	foundLogID := app.SimpleCache.GetValue(app.CACHE_LOG, *caching.GlobalCache)
	if foundLogID == nil {
		return utils.GetResponseMessage(header.AcceptLanguage,
			http.StatusUnauthorized, 8003)
	}

	app.LogID = fmt.Sprintf("%v", *foundLogID)
	app.HeaderContent = header
	return nil
}
