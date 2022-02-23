package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	if path, exists := launcher.LookPath(); exists {
		u := launcher.New().
			UserDataDir("path").
			Headless(true).
			Bin(path).MustLaunch()

		page := rod.New().ControlURL(u).MustConnect().MustPage("https://example.com")
		page.MustWaitLoad().MustScreenshot("6.png")
	}
}
