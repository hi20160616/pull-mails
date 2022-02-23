package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	path, _ := launcher.LookPath()
	// As you can see from above the -- prefix is optional
	// such as `headless` and `--headless` are the same.
	// u := launcher.New().
	//         Set("user-data-dir", "path").
	//         Set("headless").
	//         Delete("--headless").
	//         Bin(path).MustLaunch()
	u := launcher.New().
		UserDataDir("path").
		Headless(true).
		Headless(false).
		Bin(path).MustLaunch()
	page := rod.New().ControlURL(u).MustConnect().MustPage("https://example.com")
	page.MustWaitLoad().MustScreenshot("5.png")
}
