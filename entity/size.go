package entity

import (
	"github.com/google/uuid"
	"shoeShop/db"
)

type Size struct {
	Id    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Value string    `json:"value"`
}

func (_ Size) TableName() string {
	return "size"
}

func MigrateSize() {
	err := db.DB().AutoMigrate(Size{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateSize)
}
