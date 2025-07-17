package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/ocuroot/ui/components/card"
	"github.com/ocuroot/ui/components/grid"
	"github.com/ocuroot/ui/css"
)

func main() {
	flagPort := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	cssClasses := []templ.CSSClass{header()}
	cssClasses = append(cssClasses, card.AllCSS()...)
	cssClasses = append(cssClasses, grid.AllCSS()...)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		showcase := Showcase()
		showcase.Render(r.Context(), w)
	})
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		w.Write(css.Style)
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
