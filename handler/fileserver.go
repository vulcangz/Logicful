package handler

import (
	"net/http"
)

type fileHandler struct {
	root http.FileSystem
}

func FileServer(root http.FileSystem) http.Handler {
	return &fileHandler{root}
}

func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}
