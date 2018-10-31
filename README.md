# shopify-cloner

>>>
Tools to working with Shopify
>>>

## Installation

```bash
go get github.com/baorv/shopify-cloner
```

## Commands

Clone Shopify theme to local
 
```bash
shopify-cloner theme:clone -t {theme id} -a {access token} -s {shop name} -v
```

Get all available themes in store

```bash
shopify-cloner theme:list -a {access token} -s {shop name} -v
``` 

## Add new commands

Create new file in internal or any where you want

```go
package theme

import (
	"github.com/baorv/shopify-cloner/utils/parser"
)

type SomeCommand struct {
	Theme  int    `short:"t" long:"theme" description:"Theme ID you want to clone" required:"true"`
    Output string `short:"o" long:"output" description:"Output zip file you want to export"`
    Zip    bool   `short:"c" long:"compress" description:"Compress all assets to zip file"`
}
func (c *SomeCommand) Execute(args []string) error {
	return nil
}
var some SomeCommand

func init() {
	parser.Parser.AddCommand(
		"some",
		"SomeCommand description",
		"Long someCommand description",
		&some)
}

```

## Todo

## License
This project is licensed under the [MIT License](LICENSE).