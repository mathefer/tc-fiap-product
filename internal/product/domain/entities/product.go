package entities

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	Name        string    `gorm:"size:255;not null"`
	Category    int       `gorm:"not null"`
	Price       float64   `gorm:"not null"`
	Description string    `gorm:"size:255"`
	ImageLink   string    `gorm:"size:255"`
}

func (Product) TableName() string {
	return "product"
}

