package domains

import (
	"github.com/cybreem/sequis-test/app/domains/types"
	"github.com/gin-gonic/gin"
	sdkTypes "github.com/h4lim/go-sdk/app/types"
)

type FriendsManagementDelivery interface {
	FriendRequestList(c *gin.Context)
	CreateFriendRequest(c *gin.Context)
	FriendRequestAction(c *gin.Context)
	Friendlist(c *gin.Context)
	Mutuallist(c *gin.Context)
	Block(c *gin.Context)
	Unblock(c *gin.Context)
}

type FriendsManagementUseCase interface {
	FriendRequestList(param types.FriendRequestList) (*types.FriendRequestListResponse, *sdkTypes.GeneralResponse)
	CreateFriendRequest(param types.CreateFriendRequest) (*types.CreateFriendRequestResponse, *sdkTypes.GeneralResponse)
	FriendRequestAction(action string, param types.ActionFriendRequest) (*types.ActionFriendRequestResponse, *sdkTypes.GeneralResponse)
	Friendlist(param types.FriendlistRequest) (*types.FriendlistResponse, *sdkTypes.GeneralResponse)
	Mutuallist(param types.MutualRequest) (*types.MutualListResponse, *sdkTypes.GeneralResponse)
	Block(param types.BlockRequest) (*types.BlockResponse, *sdkTypes.GeneralResponse)
	Unblock(param types.BlockRequest) (*types.BlockResponse, *sdkTypes.GeneralResponse)
}

type FriendsManagementService interface {
	FriendRequestListService(types.FriendRequestList) *types.FriendRequestListResponse
	CreateFriendRequestService(types.CreateFriendRequest) (*types.CreateFriendRequestResponse, *sdkTypes.GeneralResponse)
	FriendRequestAcceptService(types.ActionFriendRequest) (*types.ActionFriendRequestResponse, *sdkTypes.GeneralResponse)
	FriendRequestRejectService(types.ActionFriendRequest) (*types.ActionFriendRequestResponse, *sdkTypes.GeneralResponse)
	FriendlistService(types.FriendlistRequest) *types.FriendlistResponse
	MutuallistService(types.MutualRequest) *types.MutualListResponse
	BlockService(types.BlockRequest) (*types.BlockResponse, *sdkTypes.GeneralResponse)
	UnblockService(types.BlockRequest) (*types.BlockResponse, *sdkTypes.GeneralResponse)
}
