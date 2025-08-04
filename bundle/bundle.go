package bundle

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

// Bundle is a collection of files that can be served as a unit.
// Used for serving CSS and JS.
type Bundle struct {
	extension string
	combined  string
	etag      string
}

// NewBundle creates an empty bundle
func NewBundle(extension string) *Bundle {
	return &Bundle{
		extension: extension,
	}
}

func (b *Bundle) Add(content []byte) {
	b.combined += "\n\n" + string(content)
	// Generate ETag from CSS content hash
	hash := sha256.Sum256([]byte(b.combined))
	b.etag = hex.EncodeToString(hash[:8]) // Use first 8 bytes for shorter ETag
}

func (s *Bundle) GetVersionedURL() string {
	return fmt.Sprintf("/%s.%s", s.etag, s.extension)
}

// Serve serves the combined content with proper HTTP headers
func (s *Bundle) Serve(w http.ResponseWriter, r *http.Request) {
	contentType := "text/plain"
	switch s.extension {
	case "css":
		contentType = "text/css; charset=utf-8"
	case "js":
		contentType = "application/javascript; charset=utf-8"
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "public, max-age=3600")
	w.Header().Set("ETag", fmt.Sprintf(`"%s"`, s.etag))

	// Check if client has cached version
	if r.Header.Get("If-None-Match") == fmt.Sprintf(`"%s"`, s.etag) {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s.combined))
}

// GetCombined returns the combined content
func (s *Bundle) GetCombined() string {
	return s.combined
}
