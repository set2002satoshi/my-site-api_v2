package controllers

import "mime/multipart"

type Context interface {
	Param(key string) string
	Bind(obj interface{}) error
	BindJSON(obj interface{}) error
	JSON(code int, obj interface{})
	Get(key string) (interface{}, bool)
	FormFile(name string) (*multipart.FileHeader, error)
}
