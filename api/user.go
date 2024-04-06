package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shoeShop/engine"
	"shoeShop/entity"
	"shoeShop/storage"
	"shoeShop/template"
)

func (h *Handler) UserCreate(ctx *engine.Context) {
	var user entity.User
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&user); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	//ctx.Print(storage.UserCreate(user))
	//resultUser := storage.UserCreate(user)
	//idStr := fmt.Sprintf("%d", resultUser.Id)
	//accessStr := fmt.Sprintf("%d", resultUser.Access)
	//component := template.User(entity.UserStr{
	//	Id:        idStr,
	//	Email:     resultUser.Email,
	//	Password:  resultUser.Password,
	//	BirthDate: resultUser.BirthDate,
	//	Access:    accessStr,
	//})
	component := template.User(storage.UserCreate(user))
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) UserRead(ctx *engine.Context) {
}

func (h *Handler) UsersRead(ctx *engine.Context) {
	resultUsers := storage.UsersRead()
	var usersStr []entity.UserStr
	for _, u := range resultUsers {
		idStr := fmt.Sprintf("%d", u.Id)
		accessStr := fmt.Sprintf("%d", u.Access)

		usersStr = append(usersStr, entity.UserStr{
			Id:        idStr,
			Email:     u.Email,
			Password:  u.Password,
			BirthDate: u.BirthDate,
			Access:    accessStr,
		})
	}
	component := template.Users(usersStr)
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) UserUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var user entity.User
	if err := decoder.Decode(&user); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.UserUpdate(user, id))
}

func (h *Handler) UserDelete(ctx *engine.Context) {
}
