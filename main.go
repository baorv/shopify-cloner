package main

import (
	_ "github.com/baorv/shopify-cloner/cmd/asset"
	_ "github.com/baorv/shopify-cloner/cmd/theme"
	"github.com/jessevdk/go-flags"
	"os"
	"github.com/baorv/shopify-cloner/utils/parser"
)

func main() {
	if _, err := parser.Parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
