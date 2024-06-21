package image

import _ "embed"

const (
	ResourceTileSize int = 16
	ToolTileSize     int = 16
)

var (
	//go:embed resource_tiles.png
	ResourceTiles_png []byte

	//go:embed tool_tiles.png
	ToolTiles_png []byte
)
