package entity

import (
	"shoeShop/db"
	"time"
)

type Token struct {
	Id      uint32    `json:"-" gorm:"primaryKey"`
	UserId  uint32    `json:"user_id"`
	Token   string    `json:"token"`
	Expired time.Time `json:"expired"`
}

func (_ Token) TableName() string {
	return "token"
}

func MigrateToken() {
	err := db.DB().AutoMigrate(Token{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateToken)
}
