package notify

import (
	"log"

	"github.com/deckarep/gosx-notifier"
)

// Notify ...
func Notify(appName string, title string, text string, iconPath string) {
	head := ""
	if text == "" {
		head = title
		title = ""
	} else {
		head = text
	}
	note := gosxnotifier.NewNotification(head)
	note.Title = appName
	note.Subtitle = title
	note.AppIcon = iconPath // (10.9+ ONLY)

	err := note.Push()
	if err != nil {
		log.Println("ERROR:", err)
	}
}
