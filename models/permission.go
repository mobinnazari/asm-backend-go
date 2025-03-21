package models

type Permission struct {
	ID     uint   `json:"id" gorm:"primarykey"`
	UserID uint   `json:"userID"`
	Target string `json:"target"`
	Access string `json:"access"`
}
