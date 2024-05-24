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

func (h *Handler) CategoryCreate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var category entity.Category
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&category); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	storage.CategoryCreate(category)
	component := template.Categories(storage.CategoriesRead())
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) CategoryRead(ctx *engine.Context) {

}

func (h *Handler) CategoriesRead(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	component := template.Categories(storage.CategoriesRead())
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) CategoryUpdate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var categoryObject entity.Category
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&categoryObject); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	id := GetIdFromContext(ctx)

	var category = entity.Category{
		Id:   categoryObject.Id,
		Name: categoryObject.Name,
	}

	component := template.Category(storage.CategoryUpdate(category, id))
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) CategoryEdit(ctx *engine.Context) {
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
	category := storage.CategoryRead(id)
	component := template.CategoryEdit(id, category)
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) CategoryDelete(ctx *engine.Context) {
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
	storage.CategoryDelete(id)
	ctx.Response.WriteHeader(200)
}

func (h *Handler) CategoryOptions(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	categories := storage.CategoriesRead()
	component := template.CategoryOptions(categories)
	component.Render(ctx.Request.Context(), ctx.Response)
}
