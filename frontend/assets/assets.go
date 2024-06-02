package assets

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var _files embed.FS

var Files, _ = fs.Sub(_files, "dist")
