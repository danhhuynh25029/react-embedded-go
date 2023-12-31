package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed _ui/front-end/build
var UI embed.FS

var uiFS fs.FS

func init() {
	var err error
	fsys := fs.FS(UI)
	uiFS, _ = fs.Sub(fsys, "_ui/front-end/build")

	if err != nil {
		log.Fatal("failed to get ui fs", err)
	}
}

func main() {
	http.Handle("/", handleStatic())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("err start server")
	}
}

func handleStatic() http.Handler {
	fsys := fs.FS(UI)
	html, _ := fs.Sub(fsys, "_ui/front-end/build")

	return http.FileServer(http.FS(html))
}
