package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetImagesScansSupportedFilesInNameOrder(t *testing.T) {
	tempDir := t.TempDir()
	photoDir := filepath.Join(tempDir, "photos")
	if err := os.MkdirAll(filepath.Join(photoDir, "nested"), 0755); err != nil {
		t.Fatal(err)
	}

	files := []string{
		"b.gif",
		"A.JPG",
		"note.txt",
		filepath.Join("nested", "c.png"),
	}
	for _, name := range files {
		path := filepath.Join(photoDir, name)
		if err := os.WriteFile(path, []byte("test"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	previousConfigDir := configDir
	previousCfg := Cfg
	defer func() {
		configDir = previousConfigDir
		Cfg = previousCfg
		setAllowedPhotos(nil)
	}()

	configDir = tempDir
	Cfg = Config{PhotoPath: "photos"}

	got := (&App{}).GetImages()
	want := []string{"/photos/A.JPG", "/photos/b.gif"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetImages() = %#v, want %#v", got, want)
	}

	allowedPhotosMu.RLock()
	defer allowedPhotosMu.RUnlock()
	if _, ok := allowedPhotos["A.JPG"]; !ok {
		t.Fatal("expected A.JPG to be allowed")
	}
	if _, ok := allowedPhotos["note.txt"]; ok {
		t.Fatal("did not expect unsupported files to be allowed")
	}
	if _, ok := allowedPhotos["c.png"]; ok {
		t.Fatal("did not expect nested files to be allowed")
	}
}
