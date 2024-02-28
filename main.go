package main

import (
	"fmt"
	"github.com/gek64/gek/gDownloader"
	"github.com/gek64/gek/gToolbox"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"redl/internal"
)

func main() {
	var github string
	var gitlab string
	var sourceforge string
	var included_parts cli.StringSlice
	var excluded_parts cli.StringSlice
	var output string

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "github",
			Aliases:     []string{"gh"},
			Usage:       "set github repo url",
			Destination: &github,
		},
		&cli.StringFlag{
			Name:        "gitlab",
			Aliases:     []string{"gl"},
			Usage:       "set gitlab repo url",
			Destination: &gitlab,
		},
		&cli.StringFlag{
			Name:        "sourceforge",
			Aliases:     []string{"sf"},
			Usage:       "set sourceforge repo url",
			Destination: &sourceforge,
		},
		&cli.StringSliceFlag{
			Name:        "included_parts",
			Aliases:     []string{"p"},
			Usage:       "set release file name included parts",
			Destination: &included_parts,
		},
		&cli.StringSliceFlag{
			Name:        "excluded_parts",
			Aliases:     []string{"ep"},
			Usage:       "set release file name excluded parts",
			Destination: &excluded_parts,
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "set output file",
			Destination: &output,
		},
	}

	// 打印版本函数
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("%s", cCtx.App.Version)
	}

	app := &cli.App{
		Usage:   "Release Download Tool",
		Version: "v2.00",
		Flags:   flags,
		Action: func(ctx *cli.Context) (err error) {
			var downloadLink string

			// 获取下载链接
			if github != "" {
				downloadLink, err = internal.GetGithubDownloadLink(github, included_parts.Value(), excluded_parts.Value())
				if err != nil {
					return err
				}
			} else if gitlab != "" {
				fmt.Println("download release from gitlab is under development")
			} else if sourceforge != "" {
				downloadLink, err = internal.GetSourceForgeDownloadLink(sourceforge, included_parts.Value(), excluded_parts.Value())
				if err != nil {
					return err
				}
			}

			// 下载
			if downloadLink != "" {
				err = gToolbox.CheckToolbox([]string{"curl"})
				if err != nil {
					err := gDownloader.Download(downloadLink, output, "")
					if err != nil {
						return err
					}
				} else {
					err := gDownloader.DownloadWithCurl(downloadLink, output, "")
					if err != nil {
						return err
					}
				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
