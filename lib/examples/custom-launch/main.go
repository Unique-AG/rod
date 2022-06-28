package main

import (
	"fmt"

	"github.com/Unique-AG/rod"
	"github.com/Unique-AG/rod/lib/launcher"
)

func main() {
	l := launcher.New()

	// For more info: https://pkg.go.dev/github.com/Unique-AG/rod/lib/launcher
	u := l.MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	page := browser.MustPage("http://example.com")

	fmt.Println(page.MustInfo().Title)
}
