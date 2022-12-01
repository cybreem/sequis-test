package deliveries

import (
	"reflect"
	"testing"

	"github.com/cybreem/sequis-test/app/domains"
	"github.com/gin-gonic/gin"
)

func TestNewMiddlewareDelivery(t *testing.T) {
	tests := []struct {
		name string
		want domains.MiddlewareDelivery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMiddlewareDelivery(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMiddlewareDelivery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_middlewareDeliveryContext_GroupRouter(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		m    middlewareDeliveryContext
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.GroupRouter(tt.args.c)
		})
	}
}
