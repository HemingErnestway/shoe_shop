package dto

import "github.com/google/uuid"

type Product struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Price     string    `json:"price"`
	Name      string    `json:"name"`
	ImageName string    `json:"image_name"`
	Size      string    `json:"size"`
	Amount    string    `json:"amount"`
	Brand     string    `json:"brand_id"`
	Category  string    `json:"category_id"`
	Color     string    `json:"color_id"`
	Gender    string    `json:"gender_id"`
}
