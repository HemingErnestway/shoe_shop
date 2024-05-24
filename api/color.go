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

func (h *Handler) ColorCreate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var gender entity.Color
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&gender); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	storage.ColorCreate(gender)
	component := template.Colors(storage.ColorsRead())
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) ColorRead(ctx *engine.Context) {

}

func (h *Handler) ColorsRead(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	component := template.Colors(storage.ColorsRead())
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) ColorUpdate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var genderObject entity.Color
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&genderObject); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	id := GetIdFromContext(ctx)

	var gender = entity.Color{
		Id:   genderObject.Id,
		Name: genderObject.Name,
	}

	component := template.Color(storage.ColorUpdate(gender, id))
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) ColorEdit(ctx *engine.Context) {
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
	gender := storage.ColorRead(id)
	component := template.ColorEdit(id, gender)
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) ColorDelete(ctx *engine.Context) {
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
	storage.ColorDelete(id)
	ctx.Response.WriteHeader(200)
}
