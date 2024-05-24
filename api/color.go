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

	var color entity.Color
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&color); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	storage.ColorCreate(color)
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

	var colorObject entity.Color
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&colorObject); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	id := GetIdFromContext(ctx)

	var color = entity.Color{
		Id:   colorObject.Id,
		Name: colorObject.Name,
	}

	component := template.Color(storage.ColorUpdate(color, id))
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
	color := storage.ColorRead(id)
	component := template.ColorEdit(id, color)
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

func (h *Handler) ColorOptions(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	colors := storage.ColorsRead()
	component := template.ColorOptions(colors)
	component.Render(ctx.Request.Context(), ctx.Response)
}
