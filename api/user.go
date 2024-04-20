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

	storage.UserCreate(user)
	component := template.Users(storage.UsersRead())
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) UserRead(ctx *engine.Context) {
}

func (h *Handler) UsersRead(ctx *engine.Context) {
	component := template.Users(storage.UsersRead())
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
	id := GetIdFromContext(ctx)
	storage.UserDelete(id)
	ctx.Response.WriteHeader(200)
}
