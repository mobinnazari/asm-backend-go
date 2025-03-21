package models

type User struct {
	ID             uint         `json:"id" gorm:"primarykey"`
	Email          string       `json:"email" gorm:"uniqueIndex;size:128"`
	Password       string       `json:"-"`
	OrganizationID uint         `json:"organizationID"`
	Notification   Notification `json:"notification"`
	Enabled        bool         `json:"enabled"`
	Locked         bool         `json:"locked"`
	OtpSecret      string       `json:"-"`
	Otp            bool         `json:"otp"`
	Role           string       `json:"role"`
	Permissions    []Permission `json:"permissions"`
}
