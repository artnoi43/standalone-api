package datamodel

type TokenInfo struct {
	Address     string `json:"address" gorm:"primaryKey"`
	IsScam      bool   `json:"isScam"`
	PendingScam bool   `json:"pendingScam"`
}