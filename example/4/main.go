package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).MustLaunch()
	page := rod.New().ControlURL(u).MustConnect().MustPage("https://example.com")
	page.MustWaitLoad().MustScreenshot("4.png")
}
