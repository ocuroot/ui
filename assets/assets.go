package assets

import _ "embed"

var (
	//go:embed logo.svg
	Logo string

	//go:embed favicon.ico
	Favicon []byte

	//go:embed anon_user.svg
	AnonUser []byte
)
