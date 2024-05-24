package storage

import (
	"shoeShop/db"
	"shoeShop/entity"
)

func CategoryCreate(category entity.Category) entity.Category {
	db.DB().Create(&category)
	return category
}

func CategoryRead(id string) entity.Category {
	var category entity.Category
	db.DB().Table(category.TableName()).Where(
		"id = ?", id).Find(&category)
	return category
}

func CategoriesRead() []entity.Category {
	var categories []entity.Category
	db.DB().Find(&categories)
	return categories
}

func CategoryUpdate(new entity.Category, id string) entity.Category {
	var current entity.Category
	db.DB().Table(current.TableName()).Where(
		"id = ?", id).Find(&current)

	updateFields(&current, &new)

	db.DB().Save(&current)
	return current
}

func CategoryDelete(id string) string {
	var category entity.Category
	db.DB().Table(category.TableName()).Where(
		"id = ?", id).Delete(&category)
	return ""
}
