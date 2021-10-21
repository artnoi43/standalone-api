package datamodel

import (
	"time"
)

type User struct {
	UUID      string    `json:"uuid"`
	Username  string    `json:"username" gorm:"primaryKey"`
	Password  []byte    `json:"-"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}
