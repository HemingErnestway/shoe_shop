package storage

import (
	"shoeShop/db"
	"shoeShop/entity"
)

func BrandCreate(brand entity.Brand) entity.Brand {
	db.DB().Create(&brand)
	return brand
}

func BrandRead(id string) entity.Brand {
	var brand entity.Brand
	db.DB().Table(brand.TableName()).Where(
		"id = ?", id).Find(&brand)
	return brand
}

func BrandsRead() []entity.Brand {
	var categories []entity.Brand
	db.DB().Find(&categories)
	return categories
}

func BrandUpdate(new entity.Brand, id string) entity.Brand {
	var current entity.Brand
	db.DB().Table(current.TableName()).Where(
		"id = ?", id).Find(&current)

	updateFields(&current, &new)

	db.DB().Save(&current)
	return current
}

func BrandDelete(id string) string {
	var brand entity.Brand
	db.DB().Table(brand.TableName()).Where(
		"id = ?", id).Delete(&brand)
	return ""
}
