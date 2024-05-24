package entity

import (
	"github.com/google/uuid"
	"shoeShop/db"
)

type Gender struct {
	Id   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name string    `json:"name"`
}

func (_ Gender) TableName() string {
	return "gender"
}

func MigrateGender() {
	err := db.DB().AutoMigrate(Gender{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateGender)
}
