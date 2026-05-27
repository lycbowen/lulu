package main

import (
	"context"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type App struct {
	ctx context.Context
}

var (
	allowedPhotos   = map[string]string{}
	allowedPhotosMu sync.RWMutex
)

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetConfig() Config {
	return Cfg
}

func (a *App) GetImages() []string {
	entries, err := os.ReadDir(resolvePhotoPath(Cfg.PhotoPath))
	if err != nil {
		setAllowedPhotos(nil)
		return []string{}
	}

	type photo struct {
		name string
		path string
	}
	photos := []photo{}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !isSupportedImage(name) {
			continue
		}
		photos = append(photos, photo{
			name: name,
			path: filepath.Join(resolvePhotoPath(Cfg.PhotoPath), name),
		})
	}

	sort.Slice(photos, func(i, j int) bool {
		return strings.ToLower(photos[i].name) < strings.ToLower(photos[j].name)
	})

	nextAllowed := map[string]string{}
	urls := make([]string, 0, len(photos))
	for _, item := range photos {
		nextAllowed[item.name] = item.path
		urls = append(urls, "/photos/"+url.PathEscape(item.name))
	}
	setAllowedPhotos(nextAllowed)

	return urls
}

func setAllowedPhotos(next map[string]string) {
	allowedPhotosMu.Lock()
	defer allowedPhotosMu.Unlock()
	allowedPhotos = next
	if allowedPhotos == nil {
		allowedPhotos = map[string]string{}
	}
}
