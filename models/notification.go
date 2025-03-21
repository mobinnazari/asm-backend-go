package models

type Notification struct {
	ID      uint `json:"id" gorm:"primarykey"`
	UserID  uint `json:"userID"`
	Active  bool `json:"active"`
	Suggest bool `json:"suggest"`
}
