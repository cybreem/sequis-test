package types

// import (
// 	"github.com/cybreem/sequis-test/app/domains/models"
// )

type FriendRequestList struct {
	Email string `json:"email" binding:"required"`
}

type FriendRequestListResponse struct {
	Requests []Requestor `json:"requests"`
}

type Requestor struct {
	Requestor string `json:"requestor"`
	Status    string `json:"status"`
}

type FriendRequest struct {
	Requestor string `json:"requestor" binding:"required"`
	To        string `json:"to" binding:"required"`
}

type CreateFriendRequest struct {
	FriendRequest
}

type ActionFriendRequest struct {
	FriendRequest
}

type FriendRequestResponse struct {
	Success bool `json:"success"`
}

type CreateFriendRequestResponse struct {
	Success bool `json:"success"`
}

type ActionFriendRequestResponse struct {
	Success bool `json:"success"`
}

type FriendlistRequest struct {
	Email string `json:"email"`
}

type FriendlistResponse struct {
	Friends []string `json:"friends"`
}

type MutualRequest struct {
	Friends []string `json:"friends"`
}

type MutualListResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

type BlockRequest struct {
	Requestor string `json:"requestor"`
	Block     string `json:"block"`
}

type BlockResponse struct {
	Success bool `json:"success"`
}
