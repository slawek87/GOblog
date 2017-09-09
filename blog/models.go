package blog

import (
	"github.com/jinzhu/gorm"
)

type Blog struct {
	gorm.Model
	Post 		string
}