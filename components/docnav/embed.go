package docnav

import (
	_ "embed"

	"github.com/ocuroot/ui/css"
)

//go:embed docnav.css
var CSS []byte

func init() {
	css.Default().Add(CSS)
}
