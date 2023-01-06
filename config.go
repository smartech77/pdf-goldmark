package pdf

import (
	"context"
	"io"
	"net/http"
	"os"

	"github.com/go-swiss/fonts"
	"github.com/yuin/goldmark/util"
)

type Config struct {
	Context context.Context

	PDF PDF

	// A source for images
	ImageFS http.FileSystem

	// All other options have sensible defaults
	Styles Styles

	// A cache for the fonts
	FontsCache fonts.Cache

	// For debugging
	TraceWriter io.Writer

	NodeRenderers util.PrioritizedSlice
}

func DefaultConfig() *Config {
	c := &Config{}
	c.Context = context.Background()
	c.ImageFS = http.FS(os.DirFS("."))
	c.Styles = DefaultStyles()

	return c
}

func (c *Config) AddDefaultNodeRenderers() {
	var nr NodeRenderer = &nodeRederFuncs{}

	c.NodeRenderers = append(c.NodeRenderers,
		util.Prioritized(nr, 1000),
	)
}
