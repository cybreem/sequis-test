package deliveries

import (
	"net/http"
	"time"

	"github.com/cybreem/sequis-test/app"
	"github.com/cybreem/sequis-test/app/domains"
	"github.com/cybreem/sequis-test/app/domains/types"
	"github.com/cybreem/sequis-test/app/usecases"
	"github.com/gin-gonic/gin"

	sdkUtils "github.com/h4lim/go-sdk/utils"
)

type friendsManagementDeliveryContext struct {
}

func NewFriendsManagementDelivery() domains.FriendsManagementDelivery {
	return friendsManagementDeliveryContext{}
}

//  FriendRequestList godoc
//	@Summary		Friend Request List
//	@Description	API for Friend Request List
//	@Tags			Friend List
//  @Param message body types.FriendRequestList true "Request friend request list"
//  @Param Accept-Language header string true "Must be EN or ID" default(EN)
//	@Success		200			{object}	types.FriendRequestListResponse	"response CodeSystemValidate and generic response"
//	@Failure		400			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Failure		500			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Router			/friends/request-list [POST]
func (s friendsManagementDeliveryContext) FriendRequestList(c *gin.Context) {
	app.Log.Debugf(app.LogID, "FriendRequestList() Process ... ")
	var param types.FriendRequestList
	if err := c.ShouldBindJSON(&param); err != nil {
		app.Log.Warningf(app.LogID, "Response FriendRequestList() Error : %v",
			sdkUtils.JsonToString(app.LogID, err.Error()))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), http.StatusBadRequest,
			err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"general_response": gr,
		})
		return
	}
	response, responseError := usecases.NewfriendsManagementUseCase().FriendRequestList(param)
	if responseError != nil {
		app.Log.Warningf(app.LogID, "Response FriendRequestList() Error ",
			sdkUtils.JsonToString(app.LogID, responseError))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), responseError.Code,
			sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, responseError.Code))
		c.JSON(responseError.HttpCode, gin.H{
			"general_response": gr,
		})
		return
	}

	app.Log.Debugf(app.LogID, "Response FriendRequestList() Success %v",
		sdkUtils.JsonToString(app.LogID, response))

	gr := sdkUtils.GeneralResponseSuccessBuild(time.Now(), http.StatusOK,
		sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, http.StatusOK))

	c.JSON(http.StatusOK, gin.H{
		"data":             response,
		"general_response": gr,
	})
}

//  FriendRequestList godoc
//	@Summary		Friend Request
//	@Description	API for Friend Request
//	@Tags			Friends Management
//  @Param message body types.CreateFriendRequest true "Request friend request"
//  @Param Accept-Language header string true "Must be EN or ID" default(EN)
//	@Success		200			{object}	types.CreateFriendRequestResponse	"response CreateFriendRequestResponse and generic response"
//	@Failure		400			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Failure		500			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Router			/friends/create-request [POST]
func (s friendsManagementDeliveryContext) CreateFriendRequest(c *gin.Context) {
	app.Log.Debugf(app.LogID, "CreateFriendRequest() Process ... ")
	var param types.CreateFriendRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		app.Log.Warningf(app.LogID, "Response CreateFriendRequest() Error : %v",
			sdkUtils.JsonToString(app.LogID, err.Error()))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), http.StatusBadRequest,
			err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"general_response": gr,
		})
		return
	}
	response, responseError := usecases.NewfriendsManagementUseCase().CreateFriendRequest(param)
	if responseError != nil {
		app.Log.Warningf(app.LogID, "Response CreateFriendRequest() Error ",
			sdkUtils.JsonToString(app.LogID, responseError))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), responseError.Code,
			sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, responseError.Code))
		c.JSON(responseError.HttpCode, gin.H{
			"general_response": gr,
		})
		return
	}

	app.Log.Debugf(app.LogID, "Response CreateFriendRequest() Success %v",
		sdkUtils.JsonToString(app.LogID, response))

	gr := sdkUtils.GeneralResponseSuccessBuild(time.Now(), http.StatusOK,
		sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, http.StatusOK))

	c.JSON(http.StatusOK, gin.H{
		"data":             response,
		"general_response": gr,
	})
}

//  FriendRequestList godoc
//	@Summary		Friend Request Action
//	@Description	API for Friend Request List
//	@Tags			Friends Management
//  @Param        action   path      string  true  "(accept or reject)"
//  @Param message body types.ActionFriendRequest true "Request friend request list"
//  @Param Accept-Language header string true "Must be EN or ID" default(EN)
//	@Success		200			{object}	types.ActionFriendRequestResponse	"response ActionFriendRequestResponse and generic response"
//	@Failure		400			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Failure		500			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Router			/friends/request-action/{action} [POST]
func (s friendsManagementDeliveryContext) FriendRequestAction(c *gin.Context) {
	app.Log.Debugf(app.LogID, "CreateFriendRequest() Process ... ")
	var param types.ActionFriendRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		app.Log.Warningf(app.LogID, "Response CreateFriendRequest() Error : %v",
			sdkUtils.JsonToString(app.LogID, err.Error()))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), http.StatusBadRequest,
			err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"general_response": gr,
		})
		return
	}
	response, responseError := usecases.NewfriendsManagementUseCase().FriendRequestAction(c.Param("action"), param)
	if responseError != nil {
		app.Log.Warningf(app.LogID, "Response CreateFriendRequest() Error ",
			sdkUtils.JsonToString(app.LogID, responseError))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), responseError.Code,
			sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, responseError.Code))
		c.JSON(responseError.HttpCode, gin.H{
			"general_response": gr,
		})
		return
	}

	app.Log.Debugf(app.LogID, "Response CreateFriendRequest() Success %v",
		sdkUtils.JsonToString(app.LogID, response))

	gr := sdkUtils.GeneralResponseSuccessBuild(time.Now(), http.StatusOK,
		sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, http.StatusOK))

	c.JSON(http.StatusOK, gin.H{
		"data":             response,
		"general_response": gr,
	})
}

//  FriendRequestList godoc
//	@Summary		Friend List
//	@Description	API for Friend List
//	@Tags			Friend List
//  @Param message body types.FriendlistRequest true "Request friend request list"
//  @Param Accept-Language header string true "Must be EN or ID" default(EN)
//	@Success		200			{object}	types.FriendlistResponse	"response FriendlistResponse and generic response"
//	@Failure		400			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Failure		500			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Router			/friends/list [POST]
func (s friendsManagementDeliveryContext) Friendlist(c *gin.Context) {
	app.Log.Debugf(app.LogID, "Friendlist() Process ... ")
	var param types.FriendlistRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		app.Log.Warningf(app.LogID, "Response Friendlist() Error : %v",
			sdkUtils.JsonToString(app.LogID, err.Error()))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), http.StatusBadRequest,
			err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"general_response": gr,
		})
		return
	}
	response, responseError := usecases.NewfriendsManagementUseCase().Friendlist(param)
	if responseError != nil {
		app.Log.Warningf(app.LogID, "Response Friendlist() Error ",
			sdkUtils.JsonToString(app.LogID, responseError))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), responseError.Code,
			sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, responseError.Code))
		c.JSON(responseError.HttpCode, gin.H{
			"general_response": gr,
		})
		return
	}

	app.Log.Debugf(app.LogID, "Response Friendlist() Success %v",
		sdkUtils.JsonToString(app.LogID, response))

	gr := sdkUtils.GeneralResponseSuccessBuild(time.Now(), http.StatusOK,
		sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, http.StatusOK))

	c.JSON(http.StatusOK, gin.H{
		"data":             response,
		"general_response": gr,
	})
}

//  FriendRequestList godoc
//	@Summary		Mutual Friend
//	@Description	API for Mutual Friend
//	@Tags			Friend List
//  @Param message body types.MutualRequest true "Request Mutual Friend lists"
//  @Param Accept-Language header string true "Must be EN or ID" default(EN)
//	@Success		200			{object}	types.MutualRequest	"response MutualRequest and generic response"
//	@Failure		400			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Failure		500			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Router			/friends/mutual [POST]
func (s friendsManagementDeliveryContext) Mutuallist(c *gin.Context) {
	app.Log.Debugf(app.LogID, "Mutuallist() Process ... ")
	var param types.MutualRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		app.Log.Warningf(app.LogID, "Response Mutuallist() Error : %v",
			sdkUtils.JsonToString(app.LogID, err.Error()))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), http.StatusBadRequest,
			err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"general_response": gr,
		})
		return
	}
	response, responseError := usecases.NewfriendsManagementUseCase().Mutuallist(param)
	if responseError != nil {
		app.Log.Warningf(app.LogID, "Response Mutuallist() Error ",
			sdkUtils.JsonToString(app.LogID, responseError))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), responseError.Code,
			sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, responseError.Code))
		c.JSON(responseError.HttpCode, gin.H{
			"general_response": gr,
		})
		return
	}

	app.Log.Debugf(app.LogID, "Response Mutuallist() Success %v",
		sdkUtils.JsonToString(app.LogID, response))

	gr := sdkUtils.GeneralResponseSuccessBuild(time.Now(), http.StatusOK,
		sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, http.StatusOK))

	c.JSON(http.StatusOK, gin.H{
		"data":             response,
		"general_response": gr,
	})
}

//  FriendRequestList godoc
//	@Summary		Block User
//	@Description	API for Block User
//	@Tags			Friends Management
//  @Param message body types.BlockRequest true "Request block user"
//  @Param Accept-Language header string true "Must be EN or ID" default(EN)
//	@Success		200			{object}	types.BlockResponse	"response UnblockResponse and generic response"
//	@Failure		400			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Failure		500			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Router			/friends/block [POST]
func (s friendsManagementDeliveryContext) Block(c *gin.Context) {
	app.Log.Debugf(app.LogID, "Block() Process ... ")
	var param types.BlockRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		app.Log.Warningf(app.LogID, "Response Block() Error : %v",
			sdkUtils.JsonToString(app.LogID, err.Error()))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), http.StatusBadRequest,
			err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"general_response": gr,
		})
		return
	}
	response, responseError := usecases.NewfriendsManagementUseCase().Block(param)
	if responseError != nil {
		app.Log.Warningf(app.LogID, "Response Block() Error ",
			sdkUtils.JsonToString(app.LogID, responseError))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), responseError.Code,
			sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, responseError.Code))
		c.JSON(responseError.HttpCode, gin.H{
			"general_response": gr,
		})
		return
	}

	app.Log.Debugf(app.LogID, "Response Block() Success %v",
		sdkUtils.JsonToString(app.LogID, response))

	gr := sdkUtils.GeneralResponseSuccessBuild(time.Now(), http.StatusOK,
		sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, http.StatusOK))

	c.JSON(http.StatusOK, gin.H{
		"data":             response,
		"general_response": gr,
	})
}

//  FriendRequestList godoc
//	@Summary		Unblock User
//	@Description	API for Unblock User
//	@Tags			Friends Management
//  @Param message body types.BlockRequest true "Request unblock user"
//  @Param Accept-Language header string true "Must be EN or ID" default(EN)
//	@Success		200			{object}	types.BlockResponse	"response UnblockResponse and generic response"
//	@Failure		400			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Failure		500			{object}	types.GeneralResponseType			"for more information about response_code please see file map_data.json"
//	@Router			/friends/unblock [POST]
func (s friendsManagementDeliveryContext) Unblock(c *gin.Context) {
	app.Log.Debugf(app.LogID, "Unblock() Process ... ")
	var param types.BlockRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		app.Log.Warningf(app.LogID, "Response Unblock() Error : %v",
			sdkUtils.JsonToString(app.LogID, err.Error()))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), http.StatusBadRequest,
			err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"general_response": gr,
		})
		return
	}
	response, responseError := usecases.NewfriendsManagementUseCase().Unblock(param)
	if responseError != nil {
		app.Log.Warningf(app.LogID, "Response Unblock() Error ",
			sdkUtils.JsonToString(app.LogID, responseError))
		gr := sdkUtils.GeneralResponseErrorBuild(time.Now(), responseError.Code,
			sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, responseError.Code))
		c.JSON(responseError.HttpCode, gin.H{
			"general_response": gr,
		})
		return
	}

	app.Log.Debugf(app.LogID, "Response Unblock() Success %v",
		sdkUtils.JsonToString(app.LogID, response))

	gr := sdkUtils.GeneralResponseSuccessBuild(time.Now(), http.StatusOK,
		sdkUtils.GeneralMessageHandler(app.HeaderContent.AcceptLanguage, http.StatusOK))

	c.JSON(http.StatusOK, gin.H{
		"data":             response,
		"general_response": gr,
	})
}
