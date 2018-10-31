package parser

import (
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Verbose bool   `short:"v" long:"verbose" description:"Enable show verbose output"`
	Shop    string `short:"s" long:"shop" description:"Store name with .myshopify.com you want to clone theme" required:"true"`
	Token   string `short:"a" long:"token" description:"Access token of the store you want to clone theme" required:"true"`
	Help    bool   `short:"h" long:"help" description:"Show help message"`
}

var Opts Options

var Parser = flags.NewParser(&Opts, flags.Default)
