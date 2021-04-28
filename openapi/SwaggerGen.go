package openapi

import (
	"embed"
)

//go:embed apidocs.swagger.json
var f embed.FS

func GetSwaggerEmbedded() *embed.FS {
	return &f
}
