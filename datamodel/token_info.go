package datamodel

import "time"

type TokenInfo struct {
	Address     string    `json:"address" gorm:"primaryKey"`
	IsScam      bool      `json:"isScam"`
	PendingScam bool      `json:"pendingScam"`
	CreatedAt   time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"-" gorm:"autoUpdateTime"`
}
