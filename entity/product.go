package entity

import (
	"shoeShop/db"
)

type Product struct {
	Id uint32 `json:"id" gorm:"primaryKey"`
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
