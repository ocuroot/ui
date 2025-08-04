package js

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

//go:embed htmx.min.js
var htmxJS string

func init() {
	Register("htmx", []byte(htmxJS))
}

// Service provides unified JavaScript serving functionality
type Service struct {
	combinedJS string
	etag       string
}

type jsRegistration struct {
	Name string
	JS   []byte
}

var registeredJS []jsRegistration

func Register(name string, content []byte) {
	registeredJS = append(registeredJS, jsRegistration{
		Name: name,
		JS:   content,
	})
}

// NewService creates a new JS service with combined JavaScript content
func NewService() *Service {
	combined := combineJS()
	// Generate ETag from JS content hash
	hash := sha256.Sum256([]byte(combined))
	etag := hex.EncodeToString(hash[:8]) // Use first 8 bytes for shorter ETag

	return &Service{
		combinedJS: combined,
		etag:       etag,
	}
}

func (s *Service) GetVersionedURL() string {
	return fmt.Sprintf("/%s.js", s.etag)
}

// combineJS concatenates the JavaScript files with proper separators
func combineJS() string {
	var builder strings.Builder

	// Add registered JS
	for _, js := range registeredJS {
		builder.WriteString("\n\n/* =========================================\n")
		builder.WriteString(" * " + js.Name + "\n")
		builder.WriteString(" * ========================================= */\n\n")
		builder.WriteString(string(js.JS))
	}

	return builder.String()
}

// ServeJS serves the combined JavaScript with proper HTTP headers
func (s *Service) ServeJS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	w.Header().Set("ETag", fmt.Sprintf(`"%s"`, s.etag))

	// Check if client has cached version
	if r.Header.Get("If-None-Match") == fmt.Sprintf(`"%s"`, s.etag) {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s.combinedJS))
}

// Render implements the templ.Component interface, allowing the JS service
// to be used directly in templates
func (s *Service) Render(ctx context.Context, w io.Writer) error {
	_, err := w.Write([]byte(s.combinedJS))
	return err
}

// GetCombinedJS returns the combined JavaScript content
func (s *Service) GetCombinedJS() string {
	return s.combinedJS
}
