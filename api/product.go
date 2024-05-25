package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shoeShop/db"
	"shoeShop/dto"
	"shoeShop/engine"
	"shoeShop/entity"
	"shoeShop/storage"
	"shoeShop/template"
	"strconv"
)

func (h *Handler) ProductCreate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var product entity.Product
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&product); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	storage.ProductCreate(product)

	products := storage.ProductsRead()
	var productCards []dto.Product

	for _, product := range products {
		var brand entity.Brand
		var category entity.Category
		var color entity.Color
		var gender entity.Gender

		db.DB().Table(brand.TableName()).Where("id = ?", product.BrandId).Find(&brand)
		db.DB().Table(category.TableName()).Where("id = ?", product.CategoryId).Find(&category)
		db.DB().Table(color.TableName()).Where("id = ?", product.ColorId).Find(&color)
		db.DB().Table(gender.TableName()).Where("id = ?", product.GenderId).Find(&gender)

		var productCard = dto.Product{
			Id:        product.Id,
			Price:     product.Price,
			Name:      product.Name,
			ImageName: product.ImageName,
			Size:      product.Size,
			Amount:    product.Amount,
			Brand:     brand.Name,
			Category:  category.Name,
			Color:     color.Name,
			Gender:    gender.Name,
		}

		productCards = append(productCards, productCard)
	}

	component := template.Products(productCards)
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) ProductRead(ctx *engine.Context) {

}

func (h *Handler) ProductsRead(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	products := storage.ProductsRead()
	var productCards []dto.Product

	for _, product := range products {
		var brand entity.Brand
		var category entity.Category
		var color entity.Color
		var gender entity.Gender

		db.DB().Table(brand.TableName()).Where("id = ?", product.BrandId).Find(&brand)
		db.DB().Table(category.TableName()).Where("id = ?", product.CategoryId).Find(&category)
		db.DB().Table(color.TableName()).Where("id = ?", product.ColorId).Find(&color)
		db.DB().Table(gender.TableName()).Where("id = ?", product.GenderId).Find(&gender)

		var productCard = dto.Product{
			Id:        product.Id,
			Price:     product.Price,
			Name:      product.Name,
			ImageName: product.ImageName,
			Size:      product.Size,
			Amount:    product.Amount,
			Brand:     brand.Name,
			Category:  category.Name,
			Color:     color.Name,
			Gender:    gender.Name,
		}

		productCards = append(productCards, productCard)
	}

	component := template.Products(productCards)
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) ProductUpdate(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	if !session.Values["authenticated"].(bool) {
		http.Error(ctx.Response, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var productObject entity.Product
	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&productObject); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}

	id := GetIdFromContext(ctx)
	productNew := storage.ProductUpdate(productObject, id)

	var brand entity.Brand
	var category entity.Category
	var color entity.Color
	var gender entity.Gender

	db.DB().Table(brand.TableName()).Where("id = ?", productNew.BrandId).Find(&brand)
	db.DB().Table(category.TableName()).Where("id = ?", productNew.CategoryId).Find(&category)
	db.DB().Table(color.TableName()).Where("id = ?", productNew.ColorId).Find(&color)
	db.DB().Table(gender.TableName()).Where("id = ?", productNew.GenderId).Find(&gender)

	var product = dto.Product{
		Id:        productNew.Id,
		Price:     productNew.Price,
		Name:      productNew.Name,
		ImageName: productNew.ImageName,
		Size:      productNew.Size,
		Amount:    productNew.Amount,
		Brand:     brand.Name,
		Category:  category.Name,
		Color:     color.Name,
		Gender:    gender.Name,
	}

	component := template.Product(product)
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) ProductEdit(ctx *engine.Context) {
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
	product := storage.ProductRead(id)
	component := template.ProductEdit(id, product)
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) ProductDelete(ctx *engine.Context) {
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
	storage.ProductDelete(id)
	ctx.Response.WriteHeader(200)
}

func (h *Handler) ProductCards(ctx *engine.Context) {
	products := storage.ProductsRead()
	var productCards []dto.Product

	for _, product := range products {
		var brand entity.Brand
		var category entity.Category
		var color entity.Color
		var gender entity.Gender

		db.DB().Table(brand.TableName()).Where("id = ?", product.BrandId).Find(&brand)
		db.DB().Table(category.TableName()).Where("id = ?", product.CategoryId).Find(&category)
		db.DB().Table(color.TableName()).Where("id = ?", product.ColorId).Find(&color)
		db.DB().Table(gender.TableName()).Where("id = ?", product.GenderId).Find(&gender)

		var productCard = dto.Product{
			Id:        product.Id,
			Price:     product.Price,
			Name:      product.Name,
			ImageName: product.ImageName,
			Size:      product.Size,
			Amount:    product.Amount,
			Brand:     brand.Name,
			Category:  category.Name,
			Color:     color.Name,
			Gender:    gender.Name,
		}

		productCards = append(productCards, productCard)
	}

	component := template.ProductCards(productCards)
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) ProductOrder(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	product := storage.ProductRead(id)

	var brand entity.Brand
	var category entity.Category
	var color entity.Color
	var gender entity.Gender

	db.DB().Table(brand.TableName()).Where("id = ?", product.BrandId).Find(&brand)
	db.DB().Table(category.TableName()).Where("id = ?", product.CategoryId).Find(&category)
	db.DB().Table(color.TableName()).Where("id = ?", product.ColorId).Find(&color)
	db.DB().Table(gender.TableName()).Where("id = ?", product.GenderId).Find(&gender)

	var productOrder = dto.Product{
		Id:        product.Id,
		Price:     product.Price,
		Name:      product.Name,
		ImageName: product.ImageName,
		Size:      product.Size,
		Amount:    product.Amount,
		Brand:     brand.Name,
		Category:  category.Name,
		Color:     color.Name,
		Gender:    gender.Name,
	}

	component := template.ProductOrder(id, productOrder)
	component.Render(ctx.Request.Context(), ctx.Response)
}

func (h *Handler) ProductOrderConfirm(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	product := storage.ProductRead(id)

	fmt.Println(product.Amount)

	amount, err := decreaseProductAmount(product.Amount)
	if err != nil {
		ctx.Response.Write([]byte("Sold out"))
		return
	}

	product.Amount = amount
	storage.ProductUpdate(product, id)

	ctx.Response.Write([]byte("Successfully purchased"))
}

func decreaseProductAmount(amount string) (string, error) {
	a, err := strconv.Atoi(amount)
	if err != nil {
		return amount, err
	}

	if a > 0 {
		a -= 1
	}
	return fmt.Sprintf("%d", a), nil
}
