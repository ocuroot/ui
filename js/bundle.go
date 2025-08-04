package js

import (
	_ "embed"

	"github.com/ocuroot/ui/bundle"
)

//go:embed htmx.min.js
var htmxJS string

var defaultBundle *bundle.Bundle

func Default() *bundle.Bundle {
	if defaultBundle == nil {
		defaultBundle = bundle.NewBundle("js")
	}
	return defaultBundle
}

func init() {
	Default().Add([]byte(htmxJS))
}
