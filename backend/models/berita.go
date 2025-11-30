package models

import "time"

type Berita struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"type:text;not null" json:"title"`
	CategoryID uint      `json:"category_id"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryID"`
	Content    string    `gorm:"type:text" json:"content"`
	Excerpt    string    `gorm:"type:text" json:"excerpt"`
	Cover      string    `gorm:"type:text" json:"cover"`
	Views      int       `gorm:"default:0" json:"views"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
