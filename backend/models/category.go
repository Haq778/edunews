package models

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);unique;not null" json:"name"`
	Slug      string    `gorm:"type:varchar(150);unique;not null" json:"slug"`
	CreatedAt time.Time `json:"created_at"`
}
