package dto

import "time"

type GetProductResponseDto struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Category    int       `json:"category"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	ImageLink   string    `json:"image_link"`
}

