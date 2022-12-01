package usecases

import (
	"net/http"

	"github.com/cybreem/sequis-test/app"
	"github.com/cybreem/sequis-test/app/domains"
	"github.com/cybreem/sequis-test/app/domains/types"
	"github.com/cybreem/sequis-test/app/services"
	sdkTypes "github.com/h4lim/go-sdk/app/types"
	sdkUtils "github.com/h4lim/go-sdk/utils"
)

type friendManagementUseCaseContext struct {
}

func NewfriendsManagementUseCase() domains.FriendsManagementUseCase {
	return friendManagementUseCaseContext{}
}

func (s friendManagementUseCaseContext) FriendRequestList(param types.FriendRequestList) (*types.FriendRequestListResponse, *sdkTypes.GeneralResponse) {
	friendManagementService := services.NewFriendsManagementServices()

	// searchParam := friendManagementService.BuildRestaurantSearchParamService(param)
	fetchFriendRequests := friendManagementService.FriendRequestListService(param)
	return fetchFriendRequests, nil
}

func (s friendManagementUseCaseContext) CreateFriendRequest(param types.CreateFriendRequest) (*types.CreateFriendRequestResponse, *sdkTypes.GeneralResponse) {
	friendManagementService := services.NewFriendsManagementServices()
	friendRequestProcess, err := friendManagementService.CreateFriendRequestService(param)
	return friendRequestProcess, err
}

func (s friendManagementUseCaseContext) FriendRequestAction(action string, param types.ActionFriendRequest) (*types.ActionFriendRequestResponse, *sdkTypes.GeneralResponse) {
	friendManagementService := services.NewFriendsManagementServices()

	var friendRequestProcess *types.ActionFriendRequestResponse
	var err *sdkTypes.GeneralResponse

	if action == "accept" {
		friendRequestProcess, err = friendManagementService.FriendRequestAcceptService(param)
	} else if action == "reject" {
		friendRequestProcess, err = friendManagementService.FriendRequestRejectService(param)
	} else {
		err := sdkUtils.GetResponseMessage(app.HeaderContent.AcceptLanguage, http.StatusBadRequest, 400)
		return nil, err
	}
	if err != nil {
		return nil, err
	} else {
		return friendRequestProcess, nil
	}
}

func (s friendManagementUseCaseContext) Friendlist(param types.FriendlistRequest) (*types.FriendlistResponse, *sdkTypes.GeneralResponse) {
	friendManagementService := services.NewFriendsManagementServices()

	fetchFriendlists := friendManagementService.FriendlistService(param)
	return fetchFriendlists, nil
}

func (s friendManagementUseCaseContext) Mutuallist(param types.MutualRequest) (*types.MutualListResponse, *sdkTypes.GeneralResponse) {
	friendManagementService := services.NewFriendsManagementServices()

	fetchMutuallists := friendManagementService.MutuallistService(param)
	return fetchMutuallists, nil
}

func (s friendManagementUseCaseContext) Block(param types.BlockRequest) (*types.BlockResponse, *sdkTypes.GeneralResponse) {
	friendManagementService := services.NewFriendsManagementServices()

	friendRequestProcess, err := friendManagementService.BlockService(param)
	if err == nil {
		return friendRequestProcess, nil
	} else {
		return nil, err
	}
}

func (s friendManagementUseCaseContext) Unblock(param types.BlockRequest) (*types.BlockResponse, *sdkTypes.GeneralResponse) {
	friendManagementService := services.NewFriendsManagementServices()

	unblockProcess, err := friendManagementService.UnblockService(param)
	if err == nil {
		return unblockProcess, nil
	} else {
		return nil, err
	}
}
