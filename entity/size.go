package entity

import "shoeShop/db"

type Size struct {
	Id    uint32 `json:"id" gorm:"primaryKey"`
	Value string `json:"value"`
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
