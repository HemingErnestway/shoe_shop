package main

import (
	"log"
	"net/http"
	"reflect"
	"shoeShop/api"
	"shoeShop/config"
	"shoeShop/db"
	"shoeShop/engine"
	"shoeShop/entity"
	"slices"
	"strings"
	"time"
)

type Info struct {
	Name   string
	Access string
	Link   reflect.Value
}

var handler *api.Handler
var staticFileFormats map[string]bool
var HttpMethodUrlFunction map[string]map[string]Info
var accessExceptions []string

func init() {
	cfg := config.Get()

	HttpMethodUrlFunction = make(map[string]map[string]Info)

	HttpMethodUrlFunction["POST"] = make(map[string]Info)
	HttpMethodUrlFunction["PUT"] = make(map[string]Info)
	HttpMethodUrlFunction["DELETE"] = make(map[string]Info)
	HttpMethodUrlFunction["GET"] = make(map[string]Info)

	maps := cfg.Api

	staticFileFormats = make(map[string]bool)

	staticFileFormats[".html"] = true
	staticFileFormats[".css"] = true
	staticFileFormats[".js"] = true
	staticFileFormats[".png"] = true
	staticFileFormats[".svg"] = true

	handler = &api.Handler{}

	services := reflect.ValueOf(handler)
	_struct := reflect.TypeOf(handler)

	for methodNum := 0; methodNum < _struct.NumMethod(); methodNum++ {
		method := _struct.Method(methodNum)
		val, ok := maps[method.Name]
		if !ok {
			continue
		}
		if _, ok := HttpMethodUrlFunction[val.Method]; !ok {

		}
		HttpMethodUrlFunction[val.Method][val.Url] = Info{
			Name:   method.Name,
			Access: "",
			Link:   services.Method(methodNum),
		}
	}

	accessExceptions = cfg.List
}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	ctx := engine.Context{
		Response: w,
		Request:  r,
	}

	url := r.URL
	path := url.Path[1:]
	pathArr := strings.Split(path, "/")
	pathName := pathArr[0]
	isEditing := false

	if pathArr[len(pathArr)-1] == "edit" {
		isEditing = true
		pathArr = pathArr[:len(pathArr)-1]
	}

	if pathArr[0] == "" {
		sendFile("./static/index.html", ctx)
		return
	}

	if staticUrl, ok := isStatic(path); ok {
		sendFile("./static/"+staticUrl, ctx)
		return
	}

	urlFunction, ok := HttpMethodUrlFunction[r.Method]
	if !ok {
		w.Write([]byte("No such method"))
	}

	if len(pathArr) > 1 {
		pathName += "/{id}"
	}

	if isEditing {
		pathName += "/edit"
	}

	if fun, ok := urlFunction[pathName]; ok {
		if slices.Contains(accessExceptions, fun.Name) || true {
			in := make([]reflect.Value, 1)
			in[0] = reflect.ValueOf(&ctx)
			fun.Link.Call(in)
		} else {
			log.Println("Forbidden")
		}
	}
}

func checkAccess(h http.Header) bool {
	authToken := h.Get("Authorization")

	log.Println("Authorization: " + authToken)

	if len(authToken) < 1 {
		return false
	}

	var token entity.Token
	db.DB().Where("token = ? and expired > ?",
		authToken, time.Now()).First(&token)

	return token.Id != 0
}

func isStatic(path string) (string, bool) {
	splitPath := strings.Split(path, "/")
	splitName := strings.Split(splitPath[len(splitPath)-1], ".")
	fileExt := "." + splitName[len(splitName)-1]

	if staticFileFormats[fileExt] {
		return path, true
	}
	return "", false
}

func sendFile(path string, ctx engine.Context) {
	http.ServeFile(ctx.Response, ctx.Request, path)
}
