package util

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func FormFile(ctx *gin.Context, key string) (multipart.File, *multipart.FileHeader, error) {
	return ctx.Request.FormFile(key)
}
