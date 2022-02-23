package main

import (
	"github.com/go-rod/rod"
)

func main() {
	page := rod.New().MustConnect().MustPage("https://wikipedia.org")
	page.MustWaitLoad().MustScreenshot("a.png")
}
