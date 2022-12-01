package services

import (
	"net/http"

	"github.com/cybreem/sequis-test/app"
	"github.com/cybreem/sequis-test/app/domains"

	"github.com/cybreem/sequis-test/app/domains/models"
	"github.com/cybreem/sequis-test/app/domains/types"

	// "github.com/cybreem/sequis-test/app/helpers"
	"github.com/cybreem/sequis-test/app/repositories"
	sdkTypes "github.com/h4lim/go-sdk/app/types"

	// "encoding/json"
	// "errors"
	// "github.com/h4lim/go-sdk/app/services"
	// sdkTypes "github.com/h4lim/go-sdk/app/types"
	sdkUtils "github.com/h4lim/go-sdk/utils"
	// "net/http"
)

type friendsManagementServiceContext struct {
	logID                    string
	FriendRequestsRepository domains.FriendRequestsRepository
	FriendlistsRepository    domains.FriendlistsRepository
}

func NewFriendsManagementServices() domains.FriendsManagementService {
	FriendRequestsRepository := repositories.NewFriendRequestsRepository()
	FriendlistsRepository := repositories.NewFriendlistsRepository()
	return friendsManagementServiceContext{
		logID:                    app.LogID,
		FriendRequestsRepository: FriendRequestsRepository,
		FriendlistsRepository:    FriendlistsRepository,
	}
}

func (svc friendsManagementServiceContext) FriendRequestListService(search types.FriendRequestList) *types.FriendRequestListResponse {

	FriendRequests := svc.FriendRequestsRepository.GetFriendRequests(search)
	if FriendRequests == nil {
		return nil
	}
	requestor := make([]types.Requestor, 0)
	for _, friendRequest := range *FriendRequests {
		friendRequestListResponse := types.Requestor{
			Requestor: friendRequest.Requestor,
			Status:    friendRequest.Status,
		}

		requestor = append(requestor, friendRequestListResponse)

	}
	return &types.FriendRequestListResponse{Requests: requestor}
}

func (svc friendsManagementServiceContext) CreateFriendRequestService(data types.CreateFriendRequest) (*types.CreateFriendRequestResponse, *sdkTypes.GeneralResponse) {
	models := models.FriendRequests{
		Email:     data.To,
		Requestor: data.Requestor,
		IsValid:   1,
	}
	checkDuplicate := svc.FriendRequestsRepository.SearchFriendRequests(models)
	if checkDuplicate != nil {
		err := sdkUtils.GetResponseMessage(app.HeaderContent.AcceptLanguage, http.StatusBadRequest, 400)
		return nil, err
	}
	checkBlock := svc.FriendlistsRepository.GetBlockedUser(models)
	if checkBlock != nil {
		err := sdkUtils.GetResponseMessage(app.HeaderContent.AcceptLanguage, http.StatusBadRequest, 8008)
		return nil, err
	}
	models.Status = "rejected"
	svc.FriendRequestsRepository.Unvalidate(models)
	models.Status = "pending"
	friendRequests := svc.FriendRequestsRepository.Save(models)
	if friendRequests == nil {
		err := sdkUtils.GetResponseMessage(app.HeaderContent.AcceptLanguage, http.StatusBadRequest, 400)
		return nil, err
	} else {
		return &types.CreateFriendRequestResponse{Success: true}, nil
	}
}

func (svc friendsManagementServiceContext) AddFriend(data types.ActionFriendRequest) *sdkTypes.GeneralResponse {
	friendList := models.Friendlists{
		User:    data.To,
		Friend:  data.Requestor,
		IsValid: 1,
		Block:   0,
	}
	Save := svc.FriendlistsRepository.Save(friendList)
	friendList.User = data.Requestor
	friendList.Friend = data.To
	Save2 := svc.FriendlistsRepository.Save(friendList)
	if Save == nil || Save2 == nil {
		return sdkUtils.GetResponseMessage(app.HeaderContent.AcceptLanguage, http.StatusBadRequest, 8011)
	} else {
		return nil
	}
}

func (svc friendsManagementServiceContext) FriendRequestAcceptService(data types.ActionFriendRequest) (*types.ActionFriendRequestResponse, *sdkTypes.GeneralResponse) {
	modelsFriendRequests := models.FriendRequests{
		Email:     data.To,
		Requestor: data.Requestor,
		Status:    "pending",
	}
	unvalidate := svc.FriendRequestsRepository.Unvalidate(modelsFriendRequests)
	if unvalidate != nil {
		modelsFriendRequests.Status = "accepted"
		modelsFriendRequests.IsValid = 1
		friendRequests := svc.FriendRequestsRepository.Save(modelsFriendRequests)
		if friendRequests == nil {
			err := sdkUtils.GetResponseMessage(app.HeaderContent.AcceptLanguage, http.StatusBadRequest, 8011)
			return nil, err
		} else {
			addFriend := svc.AddFriend(data)
			if addFriend == nil {
				return &types.ActionFriendRequestResponse{Success: true}, nil
			} else {
				err := sdkUtils.GetResponseMessage(app.HeaderContent.AcceptLanguage, http.StatusBadRequest, 8011)
				return nil, err
			}
		}
	} else {
		err := sdkUtils.GetResponseMessage(app.HeaderContent.AcceptLanguage, http.StatusBadRequest, 8011)
		return nil, err
	}
}

func (svc friendsManagementServiceContext) FriendRequestRejectService(data types.ActionFriendRequest) (*types.ActionFriendRequestResponse, *sdkTypes.GeneralResponse) {
	models := models.FriendRequests{
		Email:     data.To,
		Requestor: data.Requestor,
		Status:    "pending",
	}
	unvalidate := svc.FriendRequestsRepository.Unvalidate(models)
	if unvalidate != nil {
		models.Status = "rejected"
		models.IsValid = 1
		friendRequests := svc.FriendRequestsRepository.Save(models)
		if friendRequests == nil {
			err := sdkUtils.GetResponseMessage(app.HeaderContent.AcceptLanguage, http.StatusBadRequest, 8011)
			return nil, err
		} else {
			return &types.ActionFriendRequestResponse{Success: true}, nil
		}
	} else {
		err := sdkUtils.GetResponseMessage(app.HeaderContent.AcceptLanguage, http.StatusBadRequest, 8011)
		return nil, err
	}
}

func (svc friendsManagementServiceContext) FriendlistService(search types.FriendlistRequest) *types.FriendlistResponse {

	Friendlists := svc.FriendlistsRepository.GetFriendlists(search)
	if Friendlists == nil {
		return nil
	}
	friends := make([]string, 0)
	for _, friendlist := range *Friendlists {
		friendRow := friendlist.Friend

		friends = append(friends, friendRow)

	}
	return &types.FriendlistResponse{Friends: friends}
}

func (svc friendsManagementServiceContext) MutuallistService(search types.MutualRequest) *types.MutualListResponse {

	Mutuallists := svc.FriendlistsRepository.GetMutuallists(search)
	if Mutuallists == nil {
		return nil
	}
	friends := make([]string, 0)
	count := 0
	for _, friendlist := range *Mutuallists {
		friendRow := friendlist.Friend
		friends = append(friends, friendRow)
		count += 1
	}
	return &types.MutualListResponse{Success: true, Friends: friends, Count: count}
}

func (svc friendsManagementServiceContext) BlockService(data types.BlockRequest) (*types.BlockResponse, *sdkTypes.GeneralResponse) {
	friendList := models.Friendlists{
		User:    data.Requestor,
		Friend:  data.Block,
		IsValid: 1,
		Block:   1,
	}
	models := models.FriendRequests{
		Email:     data.Block,
		Requestor: data.Requestor,
	}

	checkBlock := svc.FriendlistsRepository.GetBlockedUser(models)
	if checkBlock != nil {
		err := sdkUtils.GetResponseMessage(app.HeaderContent.AcceptLanguage, http.StatusBadRequest, 8008)
		return nil, err
	}
	svc.FriendlistsRepository.Unvalidate(friendList)
	svc.FriendRequestsRepository.Unvalidate(models)
	models.Email = data.Requestor
	models.Requestor = data.Block
	svc.FriendRequestsRepository.Unvalidate(models)
	svc.FriendlistsRepository.Save(friendList)
	return &types.BlockResponse{Success: true}, nil
}

func (svc friendsManagementServiceContext) UnblockService(data types.BlockRequest) (*types.BlockResponse, *sdkTypes.GeneralResponse) {
	friendList := models.Friendlists{
		User:    data.Requestor,
		Friend:  data.Block,
		IsValid: 1,
		Block:   1,
	}
	models := models.FriendRequests{
		Email:     data.Block,
		Requestor: data.Requestor,
	}

	checkBlock := svc.FriendlistsRepository.GetBlockedUserStrict(models)
	if checkBlock == nil {
		err := sdkUtils.GetResponseMessage(app.HeaderContent.AcceptLanguage, http.StatusBadRequest, 400)
		return nil, err
	}
	svc.FriendlistsRepository.Unblock(friendList)
	return &types.BlockResponse{Success: true}, nil
}
