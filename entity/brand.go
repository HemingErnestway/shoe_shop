package entity

import (
	"github.com/google/uuid"
	"shoeShop/db"
)

type Brand struct {
	Id   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name string    `json:"name"`
}

func (_ Brand) TableName() string {
	return "brand"
}

func MigrateBrand() {
	err := db.DB().AutoMigrate(Brand{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateBrand)
}
