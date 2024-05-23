package api

import (
	"shoeShop/engine"
	"strings"
)

func GetIdFromContext(ctx *engine.Context) string {
	path := strings.Split(ctx.Request.URL.Path, "/")
	if len(path) > 2 && path[len(path)-1] == "edit" {
		path = path[:len(path)-1]
	}
	idString := path[len(path)-1]
	return idString
}
