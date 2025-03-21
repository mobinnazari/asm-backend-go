package models

import "time"

type Target struct {
	ID             uint      `json:"id" gorm:"primarykey"`
	OrganizationID uint      `json:"organizationID"`
	Target         string    `json:"target"`
	CreatedAt      time.Time `json:"createdAt"`
	Status         string    `json:"status"`
}
