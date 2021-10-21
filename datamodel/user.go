package datamodel

type User struct {
	UUID     string `json:"uuid"`
	Username string `json:"username" gorm:"primaryKey"`
	Password []byte `json:"-"`
}
