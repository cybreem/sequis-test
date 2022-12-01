package repositories

import (
	"reflect"
	"testing"

	"github.com/cybreem/sequis-test/app/domains"
	"github.com/cybreem/sequis-test/app/domains/models"
	"github.com/cybreem/sequis-test/app/domains/types"
)

func TestNewFriendlistsRepository(t *testing.T) {
	tests := []struct {
		name string
		want domains.FriendlistsRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFriendlistsRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFriendlistsRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendlistRepositoryContext_GetFriendlists(t *testing.T) {
	type args struct {
		param types.FriendlistRequest
	}
	tests := []struct {
		name string
		r    FriendlistRepositoryContext
		args args
		want *[]models.Friendlists
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.GetFriendlists(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendlistRepositoryContext.GetFriendlists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendlistRepositoryContext_GetMutuallists(t *testing.T) {
	type args struct {
		param types.MutualRequest
	}
	tests := []struct {
		name string
		r    FriendlistRepositoryContext
		args args
		want *[]models.Friendlists
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.GetMutuallists(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendlistRepositoryContext.GetMutuallists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendlistRepositoryContext_Save(t *testing.T) {
	type args struct {
		data models.Friendlists
	}
	tests := []struct {
		name string
		r    FriendlistRepositoryContext
		args args
		want *models.Friendlists
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Save(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendlistRepositoryContext.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendlistRepositoryContext_Block(t *testing.T) {
	type args struct {
		data models.Friendlists
	}
	tests := []struct {
		name string
		r    FriendlistRepositoryContext
		args args
		want *models.Friendlists
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Block(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendlistRepositoryContext.Block() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendlistRepositoryContext_GetBlockedUser(t *testing.T) {
	type args struct {
		model models.FriendRequests
	}
	tests := []struct {
		name string
		r    FriendlistRepositoryContext
		args args
		want *[]models.Friendlists
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.GetBlockedUser(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendlistRepositoryContext.GetBlockedUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendlistRepositoryContext_GetBlockedUserStrict(t *testing.T) {
	type args struct {
		model models.FriendRequests
	}
	tests := []struct {
		name string
		r    FriendlistRepositoryContext
		args args
		want *[]models.Friendlists
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.GetBlockedUserStrict(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendlistRepositoryContext.GetBlockedUserStrict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendlistRepositoryContext_Unblock(t *testing.T) {
	type args struct {
		data models.Friendlists
	}
	tests := []struct {
		name string
		r    FriendlistRepositoryContext
		args args
		want *models.Friendlists
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Unblock(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendlistRepositoryContext.Unblock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFriendlistRepositoryContext_Unvalidate(t *testing.T) {
	type args struct {
		data models.Friendlists
	}
	tests := []struct {
		name string
		r    FriendlistRepositoryContext
		args args
		want *models.Friendlists
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Unvalidate(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendlistRepositoryContext.Unvalidate() = %v, want %v", got, tt.want)
			}
		})
	}
}
