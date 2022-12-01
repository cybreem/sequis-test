package models

import (
	"github.com/jinzhu/gorm"
)

type Friendlists struct {
	gorm.Model
	User    string `db:"user"`
	Friend  string `db:"friend"`
	Block   uint   `db:"block"`
	IsValid uint   `db:"is_valid"`
}
