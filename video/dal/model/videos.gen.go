// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameVideos = "videos"

// Videos mapped from table <videos>
type Videos struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	PostUserID int64          `gorm:"column:post_user_id" json:"post_user_id"`
	Playurl    string         `gorm:"column:playurl" json:"playurl"`
	Coverurl   string         `gorm:"column:coverurl" json:"coverurl"`
	Title      string         `gorm:"column:title" json:"title"`
}

// TableName Videos's table name
func (*Videos) TableName() string {
	return TableNameVideos
}
