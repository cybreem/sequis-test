package usecases

import (
	"reflect"
	"testing"

	"github.com/cybreem/sequis-test/app/domains"
	"github.com/gin-gonic/gin"
	sdkTypes "github.com/h4lim/go-sdk/app/types"
)

func TestNewMiddlewareUseCase(t *testing.T) {
	type args struct {
		ginContext *gin.Context
	}
	tests := []struct {
		name string
		args args
		want domains.MiddlewareUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMiddlewareUseCase(tt.args.ginContext); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMiddlewareUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_middlewareUseCaseContext_CheckHeader(t *testing.T) {
	tests := []struct {
		name string
		m    middlewareUseCaseContext
		want *sdkTypes.GeneralResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.CheckHeader(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middlewareUseCaseContext.CheckHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
