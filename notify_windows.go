package notify

import (
	"log"

	toast "gopkg.in/toast.v1"
)

// Notify ...
func Notify(appName string, title string, text string, iconPath string) {
	notification := toast.Notification{
		AppID:   appName,
		Title:   title,
		Message: text,
		Icon:    iconPath,
	}
	err := notification.Push()
	if err != nil {
		log.Println("ERROR:", err)
	}
}
