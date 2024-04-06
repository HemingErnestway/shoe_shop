package engine

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}

type Error struct {
	Message string
}

func (c *Context) Error(status int, msg string) {
	c.Response.WriteHeader(status)
	errorMsg := Error{Message: msg}
	errorMarshal, err := json.Marshal(errorMsg)
	if err != nil {
		return
	}
	c.Response.Write(errorMarshal)
}

func (c *Context) Print(data interface{}) {
	dataMarshal, _ := json.Marshal(data)
	c.Response.Write(dataMarshal)
}

func ToStruct[T any](ctx *Context) (T, error) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var data T
	err := decoder.Decode(&data)
	return data, err
}
