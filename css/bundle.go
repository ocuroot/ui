package css

import (
	_ "embed"

	"github.com/ocuroot/ui/bundle"
)

//go:embed style.css
var mainCSS []byte

var defaultBundle *bundle.Bundle

func Default() *bundle.Bundle {
	if defaultBundle == nil {
		defaultBundle = bundle.NewBundle("css")
	}
	return defaultBundle
}

func init() {
	Default().Add([]byte(mainCSS))
}
