package datamodel

type TokenInfo struct {
	Address     string `json:"address" gorm:"primary_key"`
	IsScam      bool   `json:"isScam"`
	PendingScam bool   `json:"pendingScam"`
}