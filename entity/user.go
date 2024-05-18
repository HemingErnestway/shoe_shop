package entity

import (
	"shoeShop/db"
)

type User struct {
	Id        uint32 `json:"id" gorm:"primaryKey"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
	Access    uint32 `json:"access"`
}

func (_ User) TableName() string {
	return "user"
}

func MigrateUser() {
	err := db.DB().AutoMigrate(User{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateUser)
}
