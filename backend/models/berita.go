package models

import "time"

type Berita struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:text;not null" json:"title"`
	Category  string    `gorm:"type:varchar(100)" json:"category"`
	Content   string    `gorm:"type:text" json:"content"`
	Excerpt   string    `gorm:"type:text" json:"excerpt"`
	Cover     string    `gorm:"type:text" json:"cover"` // filename relative path
	Views     int       `gorm:"default:0" json:"views"`
	CreatedAt time.Time `json:"date"`
	UpdatedAt time.Time `json:"-"`
}
