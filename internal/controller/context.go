package controller

import "net/url"

type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	QueryParams() url.Values
	Param(name string) string
}
