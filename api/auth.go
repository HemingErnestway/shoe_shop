package api

import (
	"github.com/google/uuid"
	"shoeShop/db"
	"shoeShop/dto"
	"shoeShop/engine"
	"shoeShop/entity"
	"time"
)

func (h *Handler) Auth(ctx *engine.Context) {
	userDto, err := engine.ToStruct[dto.Auth](ctx)
	if err != nil {
		ctx.Error(400, "Bad user data")
	}

	var user entity.User
	db.DB().Table(user.TableName()).Where("email = ? and password = ?",
		userDto.Email, userDto.Password).Find(&user)

	if user.Id == 0 {
		ctx.Error(401, "Bad auth")
		return
	}

	token := entity.Token{
		UserId:  user.Id,
		Token:   uuid.NewString(),
		Expired: time.Now().Add(1 * time.Hour),
	}

	db.DB().Save(&token)
	ctx.Print(token)
}
