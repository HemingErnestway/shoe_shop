package entity

import (
	"github.com/google/uuid"
	"shoeShop/db"
)

type Product struct {
	Id         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Price      string    `json:"price"`
	Name       string    `json:"name"`
	ImageName  string    `json:"image_name"`
	Size       string    `json:"size"`
	Amount     string    `json:"amount"`
	BrandId    uuid.UUID `json:"brand_id"`
	CategoryId uuid.UUID `json:"category_id"`
	ColorId    uuid.UUID `json:"color_id"`
	GenderId   uuid.UUID `json:"gender_id"`
}

func (_ Product) TableName() string {
	return "product"
}

func MigrateProduct() {
	err := db.DB().AutoMigrate(Product{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateProduct)
}
