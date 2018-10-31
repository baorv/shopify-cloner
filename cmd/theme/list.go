package theme

import (
	"github.com/bold-commerce/go-shopify"
	"github.com/sirupsen/logrus"
	"os"
	"github.com/olekukonko/tablewriter"
	"strconv"
	"github.com/baorv/shopify-cloner/utils/parser"
	"github.com/baorv/shopify-cloner/utils"
)

type ListCommand struct {
}

func (c *ListCommand) Execute(args []string) error {
	var client = goshopify.NewClient(utils.App, parser.Opts.Shop, parser.Opts.Token)
	themes, err := client.Theme.List(nil)
	if err == nil {
		logrus.Infof("Available themes for shop %s", parser.Opts.Shop)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Role", "Theme store ID", "Previewable"})
		for _, theme := range themes {
			table.Append([]string{strconv.Itoa(theme.ID), theme.Name, theme.Role, strconv.Itoa(theme.ThemeStoreID), strconv.FormatBool(theme.Previewable)})
		}
		table.Render()
	}
	return err
}

var list ListCommand

func init() {
	parser.Parser.AddCommand(
		"theme:list",
		"List all theme from Shopify",
		"This command will list all theme from Shopify",
		&list)
}
