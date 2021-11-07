package datamodel

import "time"

type Todo struct {
	UUID      string    `json:"uuid" gorm:"primaryKey;column:uuid"`
	UserUUID  string    `json:"user" gorm:"column:user_uuid"`
	Text      string    `json:"todo" gorm:"column:text"`
	Done      bool      `json:"done" gorm:"column:done"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
}
