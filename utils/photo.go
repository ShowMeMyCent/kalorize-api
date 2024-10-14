package utils

import "mime/multipart"

type UploadedPhoto struct {
	Handler *multipart.FileHeader
	File    multipart.File
	Alias   string
}
