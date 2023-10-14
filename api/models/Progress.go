package models

import (
	"github.com/jinzhu/gorm"
)

type Progress struct {
	gorm.Model
	ProjectID   int
	UserID      int
	Description string `gorm:"type:varchar"`
}
