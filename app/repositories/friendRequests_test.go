package repositories

import (
	"reflect"
	"testing"

	"github.com/cybreem/sequis-test/app/domains"
	"github.com/cybreem/sequis-test/app/domains/models"
	"github.com/cybreem/sequis-test/app/domains/types"
)

func TestNewFriendRequestsRepository(t *testing.T) {
	tests := []struct {
		name string
		want domains.FriendRequestsRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFriendRequestsRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFriendRequestsRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendRequestsRepositoryContext_GetFriendRequests(t *testing.T) {
	type args struct {
		param types.FriendRequestList
	}
	tests := []struct {
		name string
		r    FriendRequestsRepositoryContext
		args args
		want *[]models.FriendRequests
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.GetFriendRequests(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendRequestsRepositoryContext.GetFriendRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendRequestsRepositoryContext_SearchFriendRequests(t *testing.T) {
	type args struct {
		model models.FriendRequests
	}
	tests := []struct {
		name string
		r    FriendRequestsRepositoryContext
		args args
		want *[]models.FriendRequests
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SearchFriendRequests(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendRequestsRepositoryContext.SearchFriendRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendRequestsRepositoryContext_Save(t *testing.T) {
	type args struct {
		data models.FriendRequests
	}
	tests := []struct {
		name string
		r    FriendRequestsRepositoryContext
		args args
		want *models.FriendRequests
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Save(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendRequestsRepositoryContext.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendRequestsRepositoryContext_Unvalidate(t *testing.T) {
	type args struct {
		data models.FriendRequests
	}
	tests := []struct {
		name string
		r    FriendRequestsRepositoryContext
		args args
		want *models.FriendRequests
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Unvalidate(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendRequestsRepositoryContext.Unvalidate() = %v, want %v", got, tt.want)
			}
		})
	}
}
