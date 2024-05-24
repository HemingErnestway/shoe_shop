package entity

import (
	"github.com/google/uuid"
	"shoeShop/db"
)

type Category struct {
	Id   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name string    `json:"name"`
}

func (_ Category) TableName() string {
	return "category"
}

func MigrateCategory() {
	err := db.DB().AutoMigrate(Category{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateCategory)
}
