package entity

import (
	"github.com/google/uuid"
	"shoeShop/db"
)

type User struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate string    `json:"birth_date"`
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
