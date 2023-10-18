package openapiv2

import (
	"embed"
)

//go:embed apidocs.swagger.json
var f embed.FS

func GetSwaggerEmbedded() *embed.FS {
	return &f
}