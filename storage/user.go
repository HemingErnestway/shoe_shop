package storage

import (
	"shoeShop/db"
	"shoeShop/entity"
)

func UserCreate(user entity.User) entity.User {
	db.DB().Create(&user)
	return user
}

func UserRead(id uint32) entity.User {
	var user entity.User
	db.DB().Table(user.TableName()).Where(
		"id = ?", id).Find(&user)
	return user
}

func UsersRead() []entity.User {
	var users []entity.User
	db.DB().Find(&users)
	return users
}

func UserUpdate(new entity.User, id uint32) entity.User {
	var current entity.User
	db.DB().Table(current.TableName()).Where(
		"id = ?", id).Find(&current)

	updateFields(&current, &new)

	db.DB().Save(&current)
	return current
}

func UserDelete(id uint32) string {
	var user entity.User
	db.DB().Table(user.TableName()).Where(
		"id = ?", id).Delete(&user)
	return ""
}
