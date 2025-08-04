package navbar

import (
	_ "embed"

	"github.com/ocuroot/ui/css"
	"github.com/ocuroot/ui/js"
)

//go:embed navbar.css
var CSS []byte

//go:embed navbar.js
var JS []byte

func init() {
	css.Default().Add(CSS)
	js.Default().Add(JS)
}
