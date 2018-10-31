package theme

import (
	"github.com/bold-commerce/go-shopify"
	"fmt"
	"time"
	"log"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"sync"
	"reflect"
	"github.com/baorv/shopify-cloner/utils"
	"github.com/baorv/shopify-cloner/utils/parser"
)

var SDL = []string{
	"image/gif",
	"image/jpeg",
	"application/vnd.ms-fontobject",
	"application/x-font-truetype",
	"application/font-woff",
	"image/png",
}

type CloneCommand struct {
	Theme  int    `short:"t" long:"theme" description:"Theme ID you want to clone" required:"true"`
	Output string `short:"o" long:"output" description:"Output zip file you want to export"`
	Zip    bool   `short:"c" long:"compress" description:"Compress all assets to zip file"`
}

func (c *CloneCommand) Execute(args []string) error {
	var wg = new(sync.WaitGroup)
	var items = 0
	output := "./download"
	if c.Output != "" {
		output = c.Output
	}
	var client = goshopify.NewClient(utils.App, parser.Opts.Shop, parser.Opts.Token)

	var (
		assets []goshopify.Asset
		err    error
	)
	for {
		var duration time.Duration = 0
		assets, err = client.Asset.List(c.Theme, nil)
		if err != nil && reflect.TypeOf(err).Name() == "RateLimitError" {
			duration++
			time.Sleep(duration * time.Second)
		}
		if err == nil || duration > 21 {
			break
		}
	}
	zipName := fmt.Sprintf("%d_%s_%d", time.Now().Unix(), parser.Opts.Shop, c.Theme)
	dir := fmt.Sprintf("%s/%s/", output, zipName)
	if err == nil {
		utils.CreateDir(dir)
		for _, asset := range assets {
			cloneTheme(wg, c.Theme, client, asset, dir)
			items++
		}
		wg.Wait()
		if c.Zip {
			utils.ZipDir(dir, fmt.Sprintf("%s%s.%s", dir, zipName, "zip"))
		}
		logrus.Infof("Clone %d successfully into %s", items, dir)
	}
	return nil
}

func cloneTheme(wg *sync.WaitGroup, theme int, client *goshopify.Client, asset goshopify.Asset, dir string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		var (
			sAsset *goshopify.Asset
			err    error
		)
		for {
			var duration time.Duration = 0
			sAsset, err = client.Asset.Get(theme, asset.Key)
			if err != nil && reflect.TypeOf(err).Name() == "RateLimitError" {
				duration++
				time.Sleep(duration * time.Second)
			}
			if err == nil || duration > 21 {
				break
			}
		}
		if err == nil {
			path := dir + sAsset.Key
			utils.CreateDir(filepath.Dir(path))
			if utils.InArray(sAsset.ContentType, SDL) {
				err = utils.DownloadFile(path, sAsset.PublicURL)
			} else {
				value := sAsset.Value
				err = utils.WriteFile(path, value)
			}
			if parser.Opts.Verbose {
				if err == nil {
					logrus.Infof("Cloned %s", sAsset.Key)
				} else {
					log.Printf("Cloned %s failed with message: %s", sAsset.Key, err.Error())
				}
			}
		} else {
			if parser.Opts.Verbose {
				log.Printf("Cloned %s failed with message: %s", asset.Key, err.Error())
			}
		}
	}()

}

var clone CloneCommand

func init() {
	parser.Parser.AddCommand(
		"theme:clone",
		"Clone a theme from Shopify",
		"This command will be clone all assets from Shopify theme to local device",
		&clone)
}
