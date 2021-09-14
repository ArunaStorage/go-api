package openapiv2

import "embed"

// content holds our static web server content.
//go:embed api/*
var Openapiv2specs embed.FS
