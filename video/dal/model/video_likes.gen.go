// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameVideoLikes = "video_likes"

// VideoLikes mapped from table <video_likes>
type VideoLikes struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UserID    int64          `gorm:"column:user_id" json:"user_id"`
	VideoID   int64          `gorm:"column:video_id" json:"video_id"`
}

// TableName VideoLikes's table name
func (*VideoLikes) TableName() string {
	return TableNameVideoLikes
}