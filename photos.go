package main

import (
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

var supportedImageExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
	".bmp":  true,
}

func isSupportedImage(name string) bool {
	return supportedImageExts[strings.ToLower(filepath.Ext(name))]
}

func servePhoto(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/photos/") {
		http.NotFound(w, r)
		return
	}

	name, err := url.PathUnescape(strings.TrimPrefix(r.URL.Path, "/photos/"))
	if err != nil || name == "" || name != filepath.Base(name) {
		http.NotFound(w, r)
		return
	}

	allowedPhotosMu.RLock()
	path, ok := allowedPhotos[name]
	allowedPhotosMu.RUnlock()
	if !ok {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, path)
}
