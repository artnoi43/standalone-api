package datamodel

import (
	"time"
)

type User struct {
	UUID      string    `json:"uuid" gorm:"primaryKey;column:uuid"`
	Username  string    `json:"username" gorm:"column:username"`
	Password  []byte    `json:"-" gorm:"column:password"`
	Todos     []Todo    `json:"todos" gorm:"foreignKey:UserUUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}
