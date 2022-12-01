package domains

import (
	// "time"

	"github.com/cybreem/sequis-test/app/domains/models"
	"github.com/cybreem/sequis-test/app/domains/types"
	// "github.com/gin-gonic/gin"
	// sdkTypes "github.com/h4lim/go-sdk/app/types"
	// sdkUtils "github.com/h4lim/go-sdk/utils"
)

type FriendRequestsRepository interface {
	GetFriendRequests(data types.FriendRequestList) *[]models.FriendRequests
	Save(data models.FriendRequests) *models.FriendRequests
	Unvalidate(data models.FriendRequests) *models.FriendRequests
	SearchFriendRequests(data models.FriendRequests) *[]models.FriendRequests
}

type FriendlistsRepository interface {
	GetFriendlists(data types.FriendlistRequest) *[]models.Friendlists
	Save(data models.Friendlists) *models.Friendlists
	Block(data models.Friendlists) *models.Friendlists
	GetMutuallists(data types.MutualRequest) *[]models.Friendlists
	GetBlockedUser(model models.FriendRequests) *[]models.Friendlists
	GetBlockedUserStrict(model models.FriendRequests) *[]models.Friendlists
	Unblock(data models.Friendlists) *models.Friendlists
	Unvalidate(data models.Friendlists) *models.Friendlists
}
