package models

type Organization struct {
	ID      uint     `json:"id" gorm:"primarykey"`
	Name    string   `json:"name" gorm:"uniqueIndex;size:128"`
	Targets []Target `json:"targets"`
	Limit   int      `json:"limit"`
	Users   []User   `json:"users"`
}
