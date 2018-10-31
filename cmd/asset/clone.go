package asset

import (
	"github.com/sirupsen/logrus"
	"github.com/bold-commerce/go-shopify"
	"path/filepath"
	"fmt"
	"time"
	"github.com/baorv/shopify-cloner/cmd/theme"
	"github.com/baorv/shopify-cloner/utils/parser"
	"github.com/baorv/shopify-cloner/utils"
)

type CloneCommand struct {
	Asset  string `short:"at" long:"asset" description:"Asset you want to clone" required:"true"`
	Theme  int    `short:"t" long:"theme" description:"Theme ID you want to clone" required:"true"`
	Output string `short:"o" long:"output" description:"Output zip file you want to export"`
}

func (c *CloneCommand) Execute(args []string) error {
	logrus.Info("Starting to clone asset")
	err, dir := c.clone()
	if err == nil {
		logrus.Infof("Cloned asset %d into %s", c.Asset, dir)
	} else {
		logrus.Errorf("Cloned failed with message %s", err.Error())
	}
	return nil
}

func (c *CloneCommand) clone() (error, string) {
	output := "./download"
	if c.Output != "" {
		output = c.Output
	}
	dir := fmt.Sprintf("%s/%d_%s_%d/", output, time.Now().Unix(), parser.Opts.Shop, c.Theme)
	var client = goshopify.NewClient(utils.App, parser.Opts.Shop, parser.Opts.Token)
	asset, err := client.Asset.Get(c.Theme, c.Asset)
	if err == nil {
		path := dir + asset.Key
		utils.CreateDir(filepath.Dir(path))
		if utils.InArray(asset.ContentType, theme.SDL) {
			err = utils.DownloadFile(path, asset.PublicURL)
		} else {
			value := asset.Value
			err = utils.WriteFile(path, value)
		}
		return err, path
	}
	return err, ""
}

var clone CloneCommand

func init() {
	parser.Parser.AddCommand(
		"asset:clone",
		"Clone a asset of theme from Shopify",
		"This command will be clone a asset of a theme from Shopify theme to local device",
		&clone)
}
