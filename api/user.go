package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shoeShop/dto"
	"shoeShop/engine"
	"shoeShop/entity"
	"shoeShop/storage"
	"shoeShop/template"
	"strconv"
)

func (h *Handler) UserCreate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

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
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	component := template.Users(storage.UsersRead())
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) UserUpdate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var dtoUser dto.User
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&dtoUser); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	id := GetIdFromContext(ctx)

	intAccess, err := strconv.ParseUint(dtoUser.Access, 10, 32)
	if err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	var user = entity.User{
		dtoUser.Id,
		dtoUser.Email,
		dtoUser.Password,
		dtoUser.BirthDate,
		uint32(intAccess),
	}

	component := template.User(storage.UserUpdate(user, id))
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) UserEdit(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	id := GetIdFromContext(ctx)
	user := storage.UserRead(id)
	component := template.UserEdit(id, user)
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) UserDelete(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	id := GetIdFromContext(ctx)
	storage.UserDelete(id)
	ctx.Response.WriteHeader(200)
}
