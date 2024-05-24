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

func (h *Handler) GenderCreate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var gender entity.Gender
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&gender); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	storage.GenderCreate(gender)
	component := template.Genders(storage.GendersRead())
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) GenderRead(ctx *engine.Context) {

}

func (h *Handler) GendersRead(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	component := template.Genders(storage.GendersRead())
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) GenderUpdate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var genderObject entity.Gender
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&genderObject); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	id := GetIdFromContext(ctx)

	var gender = entity.Gender{
		Id:   genderObject.Id,
		Name: genderObject.Name,
	}

	component := template.Gender(storage.GenderUpdate(gender, id))
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) GenderEdit(ctx *engine.Context) {
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
	gender := storage.GenderRead(id)
	component := template.GenderEdit(id, gender)
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) GenderDelete(ctx *engine.Context) {
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
	storage.GenderDelete(id)
	ctx.Response.WriteHeader(200)
}

func (h *Handler) GenderOptions(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	genders := storage.GendersRead()
	component := template.GenderOptions(genders)
	component.Render(ctx.Request.Context(), ctx.Response)
}
