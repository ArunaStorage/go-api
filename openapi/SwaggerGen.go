package main

import (
	"embed"
)

//go:embed api/*
var f embed.FS

func GetSwaggerEmbedded() *embed.FS {
	return &f
}
