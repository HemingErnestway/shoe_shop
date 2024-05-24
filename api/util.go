package api

import (
	"encoding/json"
	"fmt"
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

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
