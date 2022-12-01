package infrastucture

import (
	"net/http"

	"github.com/cybreem/sequis-test/app/deliveries"
	_ "github.com/cybreem/sequis-test/docs"

	"github.com/cybreem/sequis-test/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRoutes(debugMode bool) http.Handler {
	if !debugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.MaxMultipartMemory = 10000000 // maximal upload 10 mb

	router.Use(utils.CORSMiddleware())

	appsApiGroup := router.Group("/friends")
	appsApiGroup.Use(deliveries.NewMiddlewareDelivery().GroupRouter)
	{
		friendsManagementDelivery := deliveries.NewFriendsManagementDelivery()
		appsApiGroup.POST("/request-list", friendsManagementDelivery.FriendRequestList)
		appsApiGroup.POST("/create-request", friendsManagementDelivery.CreateFriendRequest)
		appsApiGroup.POST("/request-action/:action", friendsManagementDelivery.FriendRequestAction)
		appsApiGroup.POST("/list", friendsManagementDelivery.Friendlist)
		appsApiGroup.POST("/mutual", friendsManagementDelivery.Mutuallist)
		appsApiGroup.POST("/block", friendsManagementDelivery.Block)
		appsApiGroup.POST("/unblock", friendsManagementDelivery.Unblock)

	}

	// SWAGGER ROUTING
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
