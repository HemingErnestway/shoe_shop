package entity

import (
	"github.com/google/uuid"
	"shoeShop/db"
)

type Color struct {
	Id   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name string    `json:"name"`
}

func (_ Color) TableName() string {
	return "color"
}

func MigrateColor() {
	err := db.DB().AutoMigrate(Color{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateColor)
}
