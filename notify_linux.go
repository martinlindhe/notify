package notify

import (
	"os/exec"
)

// Notify displays a desktop notification
func Notify(appName string, title string, text string, iconPath string) {
	cmd := exec.Command("notify-send", "-i", iconPath, title, text)
	cmd.Run()
}
