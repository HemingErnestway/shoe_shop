package api

import (
	"github.com/gorilla/sessions"
	"net/http"
	"shoeShop/db"
	"shoeShop/dto"
	"shoeShop/engine"
	"shoeShop/entity"
)

var (
	store = sessions.NewCookieStore([]byte("secret"))
)

func (h *Handler) Login(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	authDto, err := engine.ToStruct[dto.Auth](ctx)
	if err != nil {
		ctx.Error(400, "Bad user data")
	}

	var user entity.User
	db.DB().Table(user.TableName()).Where("email = ? and password = ?",
		authDto.Email, authDto.Password).Find(&user)

	if user.Id == [16]byte{0} {
		ctx.Error(401, "Bad auth")
		return
	}

	session.Values["authenticated"] = true
	err = session.Save(ctx.Request, ctx.Response)
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx.Response.Write([]byte("Successfully logged in"))
	//http.Redirect(ctx.Response, ctx.Request, "/", http.StatusFound)
}

func (h *Handler) Logout(ctx *engine.Context) {
	session, err := store.Get(ctx.Request, "session")
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["authenticated"] = false
	err = session.Save(ctx.Request, ctx.Response)
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return
	}

	//ctx.Response.WriteHeader(200)
	http.Redirect(ctx.Response, ctx.Request, "/", http.StatusFound)
}
