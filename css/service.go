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

//go:embed navbar.css
var navbarCSS string

// Service provides unified CSS serving functionality
type Service struct {
	combinedCSS string
	etag        string
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

// combineCSS concatenates the main CSS and navbar CSS with proper separators
func combineCSS() string {
	var builder strings.Builder
	
	// Add main CSS
	builder.WriteString(mainCSS)
	
	// Add separator
	builder.WriteString("\n\n/* =========================================\n")
	builder.WriteString(" * NAVBAR CSS - Combined from navbar.css\n")
	builder.WriteString(" * ========================================= */\n\n")
	
	// Add navbar CSS
	builder.WriteString(navbarCSS)
	
	return builder.String()
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

// GetCSSSize returns the size of the combined CSS in bytes
func (s *Service) GetCSSSize() int {
	return len(s.combinedCSS)
}

// GetMainCSSSize returns the size of the main CSS in bytes
func (s *Service) GetMainCSSSize() int {
	return len(mainCSS)
}

// GetNavbarCSSSize returns the size of the navbar CSS in bytes
func (s *Service) GetNavbarCSSSize() int {
	return len(navbarCSS)
}
