package notify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNotify(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	Notify("app name", "title", "subtitle", filepath.Join(dir, "assets/icon128.png"))
}
