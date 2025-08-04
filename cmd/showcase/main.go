package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/ocuroot/ui/assets"
	"github.com/ocuroot/ui/css"
	"github.com/ocuroot/ui/js"

	"github.com/NYTimes/gziphandler"
)

func main() {
	flagPort := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	// Initialize the unified CSS and JS services
	cssService := css.NewService()
	jsService := js.NewService()

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
	http.HandleFunc(cssService.GetVersionedURL(), cssService.ServeCSS)
	http.HandleFunc(jsService.GetVersionedURL(), jsService.ServeJS)

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
		Handler: gziphandler.GzipHandler(http.DefaultServeMux),
	}
	fmt.Printf("Listening on port %d\n", *flagPort)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
