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

func (h *Handler) SizeCreate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var size entity.Size
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&size); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	storage.SizeCreate(size)
	component := template.Sizes(storage.SizesRead())
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) SizeRead(ctx *engine.Context) {

}

func (h *Handler) SizesRead(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	component := template.Sizes(storage.SizesRead())
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) SizeUpdate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var sizeObject entity.Size
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&sizeObject); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	id := GetIdFromContext(ctx)

	var size = entity.Size{
		Id:    sizeObject.Id,
		Value: sizeObject.Value,
	}

	component := template.Size(storage.SizeUpdate(size, id))
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) SizeEdit(ctx *engine.Context) {
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
	size := storage.SizeRead(id)
	component := template.SizeEdit(id, size)
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) SizeDelete(ctx *engine.Context) {
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
	storage.SizeDelete(id)
	ctx.Response.WriteHeader(200)
}
