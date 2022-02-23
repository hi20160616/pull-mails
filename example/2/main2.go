package main

import (
	"github.com/go-rod/rod"
)

func main() {
	u := "ws://127.0.0.1:9222/devtools/browser/16d9d9ab-e76b-4a81-9531-2a78ef79a2a9"
	page := rod.New().ControlURL(u).MustConnect().MustPage("https://example.com")
	page.MustWaitLoad().MustScreenshot("b.png")
}
