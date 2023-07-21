package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TodoEntity struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Title       string     `gorm:"not null" json:"title"`
	Completed   bool       `gorm:"not null,default:false" json:"completed"`
	Description string     `gorm:"default:''" json:"description"`
	UserID      uint       `gorm:"foreignkey:UserID" json:"user_id"`
	User        UserEntity `gorm:"foreignkey:UserID" json:"-"`
	CreatedAt   time.Time  `gorm:"" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"" json:""`
}

func (entity *TodoEntity) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	entity.UpdatedAt = time.Now().Local()
	return nil
}

func (entity *TodoEntity) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
