package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"` // hashed
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Role      string    `gorm:"type:varchar(30);default:'user'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
