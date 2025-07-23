package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/ocuroot/ui/assets"
	"github.com/ocuroot/ui/components/card"
	"github.com/ocuroot/ui/css"
	"github.com/ocuroot/ui/components/grid"
)

func main() {
	flagPort := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	// Initialize the unified CSS service
	cssService := css.NewService()
	cssClasses := []templ.CSSClass{header()}
	cssClasses = append(cssClasses, card.AllCSS()...)
	cssClasses = append(cssClasses, grid.AllCSS()...)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		showcase := Showcase()
		showcase.Render(r.Context(), w)
	})
	http.HandleFunc("/modal", func(w http.ResponseWriter, r *http.Request) {
		modalShowcase := ModalShowcase()
		modalShowcase.Render(r.Context(), w)
	})
	http.HandleFunc("/table", func(w http.ResponseWriter, r *http.Request) {
		tableShowcase := TableShowcase()
		tableShowcase.Render(r.Context(), w)
	})
	http.HandleFunc("/components", func(w http.ResponseWriter, r *http.Request) {
		componentsShowcase := ComponentsShowcase()
		componentsShowcase.Render(r.Context(), w)
	})
	http.HandleFunc("/style.css", cssService.ServeCSS)
	
	// Legacy endpoints for backward compatibility (redirect to unified CSS)
	http.HandleFunc("/css/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/style.css", http.StatusMovedPermanently)
	})
	
	http.HandleFunc("/components/navbar/navbar.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		http.ServeFile(w, r, "components/navbar/navbar.js")
	})
	http.HandleFunc("/static/js/htmx.min.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		// Minimal htmx stub to prevent errors - our navbar doesn't need htmx
		w.Write([]byte(`// Minimal HTMX stub
window.htmx = window.htmx || {};
htmx.process = htmx.process || function() {};
console.log('HTMX stub loaded');`))
	})
	http.HandleFunc("/static/logo.svg", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write([]byte(assets.Logo))
	})
	http.HandleFunc("/static/anon_user.svg", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write(assets.AnonUser)
	})
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write(assets.Favicon)
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *flagPort),
		Handler: templ.NewCSSMiddleware(http.DefaultServeMux, cssClasses...),
	}
	fmt.Printf("Listening on port %d\n", *flagPort)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
