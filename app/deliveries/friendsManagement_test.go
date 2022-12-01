package deliveries

import (
	"reflect"
	"testing"

	"github.com/cybreem/sequis-test/app/domains"
	"github.com/gin-gonic/gin"
)

func TestNewFriendsManagementDelivery(t *testing.T) {
	tests := []struct {
		name string
		want domains.FriendsManagementDelivery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFriendsManagementDelivery(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFriendsManagementDelivery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_friendsManagementDeliveryContext_FriendRequestList(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		s    friendsManagementDeliveryContext
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.FriendRequestList(tt.args.c)
		})
	}
}

func Test_friendsManagementDeliveryContext_CreateFriendRequest(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		s    friendsManagementDeliveryContext
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.CreateFriendRequest(tt.args.c)
		})
	}
}

func Test_friendsManagementDeliveryContext_FriendRequestAction(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		s    friendsManagementDeliveryContext
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.FriendRequestAction(tt.args.c)
		})
	}
}

func Test_friendsManagementDeliveryContext_Friendlist(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		s    friendsManagementDeliveryContext
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Friendlist(tt.args.c)
		})
	}
}

func Test_friendsManagementDeliveryContext_Mutuallist(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		s    friendsManagementDeliveryContext
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Mutuallist(tt.args.c)
		})
	}
}

func Test_friendsManagementDeliveryContext_Block(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		s    friendsManagementDeliveryContext
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Block(tt.args.c)
		})
	}
}

func Test_friendsManagementDeliveryContext_Unblock(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		s    friendsManagementDeliveryContext
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Unblock(tt.args.c)
		})
	}
}
