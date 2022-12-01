package usecases

import (
	"reflect"
	"testing"

	"github.com/cybreem/sequis-test/app/domains"
	"github.com/cybreem/sequis-test/app/domains/types"
	sdkTypes "github.com/h4lim/go-sdk/app/types"
)

func TestNewfriendsManagementUseCase(t *testing.T) {
	tests := []struct {
		name string
		want domains.FriendsManagementUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewfriendsManagementUseCase(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewfriendsManagementUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_friendManagementUseCaseContext_FriendRequestList(t *testing.T) {
	type args struct {
		param types.FriendRequestList
	}
	tests := []struct {
		name  string
		s     friendManagementUseCaseContext
		args  args
		want  *types.FriendRequestListResponse
		want1 *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.FriendRequestList(tt.args.param)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendManagementUseCaseContext.FriendRequestList() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("friendManagementUseCaseContext.FriendRequestList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_friendManagementUseCaseContext_CreateFriendRequest(t *testing.T) {
	type args struct {
		param types.CreateFriendRequest
	}
	tests := []struct {
		name  string
		s     friendManagementUseCaseContext
		args  args
		want  *types.CreateFriendRequestResponse
		want1 *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateFriendRequest(tt.args.param)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendManagementUseCaseContext.CreateFriendRequest() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("friendManagementUseCaseContext.CreateFriendRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_friendManagementUseCaseContext_FriendRequestAction(t *testing.T) {
	type args struct {
		action string
		param  types.ActionFriendRequest
	}
	tests := []struct {
		name  string
		s     friendManagementUseCaseContext
		args  args
		want  *types.ActionFriendRequestResponse
		want1 *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.FriendRequestAction(tt.args.action, tt.args.param)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendManagementUseCaseContext.FriendRequestAction() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("friendManagementUseCaseContext.FriendRequestAction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_friendManagementUseCaseContext_Friendlist(t *testing.T) {
	type args struct {
		param types.FriendlistRequest
	}
	tests := []struct {
		name  string
		s     friendManagementUseCaseContext
		args  args
		want  *types.FriendlistResponse
		want1 *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Friendlist(tt.args.param)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendManagementUseCaseContext.Friendlist() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("friendManagementUseCaseContext.Friendlist() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_friendManagementUseCaseContext_Mutuallist(t *testing.T) {
	type args struct {
		param types.MutualRequest
	}
	tests := []struct {
		name  string
		s     friendManagementUseCaseContext
		args  args
		want  *types.MutualListResponse
		want1 *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Mutuallist(tt.args.param)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendManagementUseCaseContext.Mutuallist() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("friendManagementUseCaseContext.Mutuallist() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_friendManagementUseCaseContext_Block(t *testing.T) {
	type args struct {
		param types.BlockRequest
	}
	tests := []struct {
		name  string
		s     friendManagementUseCaseContext
		args  args
		want  *types.BlockResponse
		want1 *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Block(tt.args.param)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendManagementUseCaseContext.Block() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("friendManagementUseCaseContext.Block() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_friendManagementUseCaseContext_Unblock(t *testing.T) {
	type args struct {
		param types.BlockRequest
	}
	tests := []struct {
		name  string
		s     friendManagementUseCaseContext
		args  args
		want  *types.BlockResponse
		want1 *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Unblock(tt.args.param)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("friendManagementUseCaseContext.Unblock() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("friendManagementUseCaseContext.Unblock() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
