package image

import _ "embed"

var (
	//go:embed resource_tiles.png
	ResourceTiles_png []byte

	//go:embed tool_tiles.png
	ToolTiles_png []byte
)
