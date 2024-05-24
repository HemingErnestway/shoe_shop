package storage

import (
	"shoeShop/db"
	"shoeShop/entity"
)

func ColorCreate(color entity.Color) entity.Color {
	db.DB().Create(&color)
	return color
}

func ColorRead(id string) entity.Color {
	var color entity.Color
	db.DB().Table(color.TableName()).Where(
		"id = ?", id).Find(&color)
	return color
}

func ColorsRead() []entity.Color {
	var categories []entity.Color
	db.DB().Find(&categories)
	return categories
}

func ColorUpdate(new entity.Color, id string) entity.Color {
	var current entity.Color
	db.DB().Table(current.TableName()).Where(
		"id = ?", id).Find(&current)

	updateFields(&current, &new)

	db.DB().Save(&current)
	return current
}

func ColorDelete(id string) string {
	var color entity.Color
	db.DB().Table(color.TableName()).Where(
		"id = ?", id).Delete(&color)
	return ""
}
