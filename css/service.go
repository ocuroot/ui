package css

import (
	"context"
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
)

//go:embed style.css
var mainCSS string

func init() {
	Register("style", []byte(mainCSS))
}

// Service provides unified CSS serving functionality
type Service struct {
	combinedCSS string
	etag        string
}

type cssRegistration struct {
	Name string
	CSS  []byte
}

var registeredCSS []cssRegistration

func Register(name string, content []byte) {
	registeredCSS = append(registeredCSS, cssRegistration{
		Name: name,
		CSS:  content,
	})
}

// NewService creates a new CSS service with combined CSS content
func NewService() *Service {
	combined := combineCSS()
	// Generate ETag from CSS content hash
	hash := sha256.Sum256([]byte(combined))
	etag := hex.EncodeToString(hash[:8]) // Use first 8 bytes for shorter ETag

	return &Service{
		combinedCSS: combined,
		etag:        etag,
	}
}

// combineCSS concatenates the main CSS, navbar CSS, and section CSS with proper separators
func combineCSS() string {
	var builder strings.Builder

	// Add registered CSS
	for _, css := range registeredCSS {
		builder.WriteString("\n\n/* =========================================\n")
		builder.WriteString(" * " + css.Name + "\n")
		builder.WriteString(" * ========================================= */\n\n")
		builder.WriteString(string(css.CSS))
	}

	return builder.String()
}

func (s *Service) GetVersionedURL() string {
	return fmt.Sprintf("/%s.css", s.etag)
}

// ServeCSS serves the combined CSS with proper HTTP headers
func (s *Service) ServeCSS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	w.Header().Set("ETag", fmt.Sprintf(`"%s"`, s.etag))

	// Check if client has cached version
	if r.Header.Get("If-None-Match") == fmt.Sprintf(`"%s"`, s.etag) {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s.combinedCSS))
}

// Render implements the templ.Component interface, allowing the CSS service
// to be used directly in templates
func (s *Service) Render(ctx context.Context, w io.Writer) error {
	_, err := w.Write([]byte(s.combinedCSS))
	return err
}

// GetCombinedCSS returns the combined CSS content
func (s *Service) GetCombinedCSS() string {
	return s.combinedCSS
}
