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

func (h *Handler) BrandCreate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var brand entity.Brand
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&brand); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	storage.BrandCreate(brand)
	component := template.Brands(storage.BrandsRead())
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) BrandRead(ctx *engine.Context) {

}

func (h *Handler) BrandsRead(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	component := template.Brands(storage.BrandsRead())
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) BrandUpdate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var brandObject entity.Brand
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&brandObject); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	id := GetIdFromContext(ctx)

	var brand = entity.Brand{
		Id:   brandObject.Id,
		Name: brandObject.Name,
	}

	component := template.Brand(storage.BrandUpdate(brand, id))
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) BrandEdit(ctx *engine.Context) {
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
	brand := storage.BrandRead(id)
	component := template.BrandEdit(id, brand)
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) BrandDelete(ctx *engine.Context) {
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
	storage.BrandDelete(id)
	ctx.Response.WriteHeader(200)
}
