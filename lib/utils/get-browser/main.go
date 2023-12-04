// Package main ...
package main

import (
	"fmt"

	"github.com/Unique-AG/rod/lib/launcher"
	"github.com/Unique-AG/rod/lib/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}
