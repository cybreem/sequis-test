package models

import (
	"github.com/jinzhu/gorm"
)

type FriendRequests struct {
	gorm.Model
	Email     string `db:"email"`
	Requestor string `db:"requestor"`
	Status    string `db:"status"`
	IsValid   uint   `db:"is_valid"`
}
