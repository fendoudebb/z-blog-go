package assets

import (
	"embed"
	"io/fs"
)

//go:embed *
var assetsFS embed.FS

func FS() fs.FS {
	return assetsFS
}
