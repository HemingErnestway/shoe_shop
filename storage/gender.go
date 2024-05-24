package storage

import (
	"shoeShop/db"
	"shoeShop/entity"
)

func GenderCreate(gender entity.Gender) entity.Gender {
	db.DB().Create(&gender)
	return gender
}

func GenderRead(id string) entity.Gender {
	var gender entity.Gender
	db.DB().Table(gender.TableName()).Where(
		"id = ?", id).Find(&gender)
	return gender
}

func GendersRead() []entity.Gender {
	var categories []entity.Gender
	db.DB().Find(&categories)
	return categories
}

func GenderUpdate(new entity.Gender, id string) entity.Gender {
	var current entity.Gender
	db.DB().Table(current.TableName()).Where(
		"id = ?", id).Find(&current)

	updateFields(&current, &new)

	db.DB().Save(&current)
	return current
}

func GenderDelete(id string) string {
	var gender entity.Gender
	db.DB().Table(gender.TableName()).Where(
		"id = ?", id).Delete(&gender)
	return ""
}
