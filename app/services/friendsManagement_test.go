package services

import (
	"reflect"
	"testing"

	"github.com/cybreem/sequis-test/app/domains"
	"github.com/cybreem/sequis-test/app/domains/types"
	sdkTypes "github.com/h4lim/go-sdk/app/types"
)

func TestNewFriendsManagementServices(t *testing.T) {
	tests := []struct {
		name string
		want domains.FriendsManagementService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFriendsManagementServices(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFriendsManagementServices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_friendsManagementServiceContext_FriendRequestListService(t *testing.T) {
	type args struct {
		search types.FriendRequestList
	}
	tests := []struct {
		name string
		svc  friendsManagementServiceContext
		args args
		want *types.FriendRequestListResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.svc.FriendRequestListService(tt.args.search); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendsManagementServiceContext.FriendRequestListService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_friendsManagementServiceContext_CreateFriendRequestService(t *testing.T) {
	type args struct {
		data types.CreateFriendRequest
	}
	tests := []struct {
		name  string
		svc   friendsManagementServiceContext
		args  args
		want  *types.CreateFriendRequestResponse
		want1 *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.svc.CreateFriendRequestService(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendsManagementServiceContext.CreateFriendRequestService() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("friendsManagementServiceContext.CreateFriendRequestService() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_friendsManagementServiceContext_AddFriend(t *testing.T) {
	type args struct {
		data types.ActionFriendRequest
	}
	tests := []struct {
		name string
		svc  friendsManagementServiceContext
		args args
		want *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.svc.AddFriend(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendsManagementServiceContext.AddFriend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_friendsManagementServiceContext_FriendRequestAcceptService(t *testing.T) {
	type args struct {
		data types.ActionFriendRequest
	}
	tests := []struct {
		name  string
		svc   friendsManagementServiceContext
		args  args
		want  *types.ActionFriendRequestResponse
		want1 *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.svc.FriendRequestAcceptService(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendsManagementServiceContext.FriendRequestAcceptService() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("friendsManagementServiceContext.FriendRequestAcceptService() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_friendsManagementServiceContext_FriendRequestRejectService(t *testing.T) {
	type args struct {
		data types.ActionFriendRequest
	}
	tests := []struct {
		name  string
		svc   friendsManagementServiceContext
		args  args
		want  *types.ActionFriendRequestResponse
		want1 *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.svc.FriendRequestRejectService(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendsManagementServiceContext.FriendRequestRejectService() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("friendsManagementServiceContext.FriendRequestRejectService() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_friendsManagementServiceContext_FriendlistService(t *testing.T) {
	type args struct {
		search types.FriendlistRequest
	}
	tests := []struct {
		name string
		svc  friendsManagementServiceContext
		args args
		want *types.FriendlistResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.svc.FriendlistService(tt.args.search); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendsManagementServiceContext.FriendlistService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_friendsManagementServiceContext_MutuallistService(t *testing.T) {
	type args struct {
		search types.MutualRequest
	}
	tests := []struct {
		name string
		svc  friendsManagementServiceContext
		args args
		want *types.MutualListResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.svc.MutuallistService(tt.args.search); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendsManagementServiceContext.MutuallistService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_friendsManagementServiceContext_BlockService(t *testing.T) {
	type args struct {
		data types.BlockRequest
	}
	tests := []struct {
		name  string
		svc   friendsManagementServiceContext
		args  args
		want  *types.BlockResponse
		want1 *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.svc.BlockService(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendsManagementServiceContext.BlockService() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("friendsManagementServiceContext.BlockService() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_friendsManagementServiceContext_UnblockService(t *testing.T) {
	type args struct {
		data types.BlockRequest
	}
	tests := []struct {
		name  string
		svc   friendsManagementServiceContext
		args  args
		want  *types.BlockResponse
		want1 *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.svc.UnblockService(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendsManagementServiceContext.UnblockService() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("friendsManagementServiceContext.UnblockService() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
