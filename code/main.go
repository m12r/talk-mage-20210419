package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed web/dist
var rootFS embed.FS

func main() {
	webFS, err := fs.Sub(rootFS, "web/dist")
	if err != nil {
		log.Fatalf("cannot mount webFS: %v", err)
	}

	http.Handle("/", http.FileServer(http.FS(webFS)))

	if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
		log.Fatalf("cannot start http server: %v", err)
	}
}
