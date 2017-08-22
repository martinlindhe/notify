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
	Notify("app name", "notify", "text", filepath.Join(dir, "assets/icon128.png"))
}

func TestAlert(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	Alert("app name", "alert!", "text", filepath.Join(dir, "assets/icon128.png"))
}
