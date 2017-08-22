# About

[![GoDoc](https://godoc.org/github.com/martinlindhe/notify?status.svg)](https://godoc.org/github.com/martinlindhe/notify)

Notify is a simple cross-platform library for displaying desktop notifications in your go application


## Example

```go
package main

import "github.com/martinlindhe/notify"

func main() {
	notify.Notify("app name", "title", "subtitle", "path/to/icon.png"))
}
```


### Windows 10

Uses Toast notficiations provided by https://github.com/go-toast/toast

![Windows](windows.png)


### macOS / OSX 10.8+

Uses terminal-notifier provided by https://github.com/deckarep/gosx-notifier

![macOS](macos.png)


### Linux

Uses the notify-send command (Gnome, KDE etc)

![Linux](linux.png)


### More

If you like this, check out https://github.com/martinlindhe/inputbox for cross-platform dialog input boxes.


### License

Under [MIT](LICENSE)
