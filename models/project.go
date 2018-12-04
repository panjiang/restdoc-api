package models

import "time"

// Project 项目
type Project struct {
	ID        int64     `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"size:30;not null" json:"name"`
	UID       int64     `gorm:"not null" json:"uid"`
	CreatedAt time.Time `json:"createdAt"`
}
