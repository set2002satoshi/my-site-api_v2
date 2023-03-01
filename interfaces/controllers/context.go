package controllers

import "mime/multipart"

type Context interface {
	
	FormFile(name string) (*multipart.FileHeader, error)
}
