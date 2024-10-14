package models

import (
	"time"

	"gorm.io/gorm"
)

// Person struct embeds gorm.Model to utilize its built-in fields
type Person struct {
	PersonID  uint           `gorm:"primaryKey;autoIncrement" json:"person_id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Age       int            `gorm:"not null" json:"age"`
	Hobbies   string         `gorm:"type:varchar(255)" json:"hobbies"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"autoUpdateTime" json:"deleted_at"`
}
