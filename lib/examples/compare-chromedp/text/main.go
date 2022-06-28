package main

import (
	"log"
	"strings"

	"github.com/Unique-AG/rod"
)

// This example demonstrates  how to extract text from a specific element.
func main() {
	page := rod.New().MustConnect().MustPage("https://golang.org/pkg/time")

	res := page.MustElement("#pkg-overview").MustText()
	log.Println(strings.TrimSpace(res))
}
