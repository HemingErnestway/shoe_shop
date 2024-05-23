package storage

import (
	"shoeShop/db"
	"shoeShop/entity"
)

func SizeCreate(size entity.Size) entity.Size {
	db.DB().Create(&size)
	return size
}

func SizeRead(id uint32) entity.Size {
	var size entity.Size
	db.DB().Table(size.TableName()).Where(
		"id = ?", id).Find(&size)
	return size
}

func SizesRead() []entity.Size {
	var sizes []entity.Size
	db.DB().Find(&sizes)
	return sizes
}

func SizeUpdate(new entity.Size, id uint32) entity.Size {
	var current entity.Size
	db.DB().Table(current.TableName()).Where(
		"id = ?", id).Find(&current)

	updateFields(&current, &new)

	db.DB().Save(&current)
	return current
}

func SizeDelete(id uint32) string {
	var size entity.Size
	db.DB().Table(size.TableName()).Where(
		"id = ?", id).Delete(&size)
	return ""
}
