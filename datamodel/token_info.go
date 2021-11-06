package datamodel

import "time"

type TokenInfo struct {
	Address     string    `json:"address" gorm:"primaryKey;column:address"`
	Chain       string    `json:"chain" gorm:"column:chain"`
	IsScam      bool      `json:"isScam" gorm:"column:is_scam"`
	PendingScam bool      `json:"pendingScam" gorm:"column:pending_scam"`
	CreatedAt   time.Time `json:"-" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"autoUpdateTime;column:updated_at"`
}
