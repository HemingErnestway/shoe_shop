package api

import (
	"shoeShop/engine"
	"strconv"
	"strings"
)

func GetIdFromContext(ctx *engine.Context) uint32 {
	path := strings.Split(ctx.Request.URL.Path, "/")
	if len(path) > 2 && path[len(path)-1] == "edit" {
		path = path[:len(path)-1]
	}
	idString := path[len(path)-1]
	idUint64, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		ctx.Error(400, err.Error())
	}
	return uint32(idUint64)
}
