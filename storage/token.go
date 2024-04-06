package storage

import (
	"shoeShop/db"
	"shoeShop/entity"
)

func TokenCreate(token entity.Token) *entity.Token {
	db.DB().Create(&token)
	return &token
}

func TokenRead(id uint32) *entity.Token {
	var token entity.Token
	db.DB().Table(token.TableName()).Where(
		"id = ?", id).Find(&token)
	return &token
}

func TokensRead() []*entity.Token {
	var tokens []*entity.Token
	db.DB().Find(&tokens)
	return tokens
}

func TokenUpdate(new entity.Token, id uint32) *entity.Token {
	var current entity.Token
	db.DB().Table(current.TableName()).Where(
		"id = ?", id).Find(&current)

	updateFields(&current, &new)

	db.DB().Save(&current)
	return &current
}

func TokenDelete(id uint32) string {
	var token entity.Token
	db.DB().Table(token.TableName()).Where(
		"id = ?", id).Delete(&token)
	return "successfully deleted"
}
