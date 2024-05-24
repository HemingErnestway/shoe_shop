package storage

import (
	"shoeShop/db"
	"shoeShop/entity"
)

func ProductCreate(product entity.Product) entity.Product {
	db.DB().Create(&product)
	return product
}

func ProductRead(id string) entity.Product {
	var product entity.Product
	db.DB().Table(product.TableName()).Where(
		"id = ?", id).Find(&product)
	return product
}

func ProductsRead() []entity.Product {
	var categories []entity.Product
	db.DB().Find(&categories)
	return categories
}

func ProductUpdate(new entity.Product, id string) entity.Product {
	var current entity.Product
	db.DB().Table(current.TableName()).Where(
		"id = ?", id).Find(&current)

	updateFields(&current, &new)

	db.DB().Save(&current)
	return current
}

func ProductDelete(id string) string {
	var product entity.Product
	db.DB().Table(product.TableName()).Where(
		"id = ?", id).Delete(&product)
	return ""
}
