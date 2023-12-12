package main

import (
	"fmt"
	"github.com/gek64/gek/gDownloader"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"redl/internal"
)

func main() {
	var github_repository string
	var sourceforge_rss string
	var included_parts cli.StringSlice
	var excluded_parts cli.StringSlice
	var output string

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "github_repository",
			Aliases:     []string{"gr"},
			Usage:       "set github repository",
			Destination: &github_repository,
		},
		&cli.StringFlag{
			Name:        "sourceforge_rss",
			Aliases:     []string{"sr"},
			Usage:       "set sourceforge rss",
			Destination: &sourceforge_rss,
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
		Action: func(ctx *cli.Context) error {
			var downloadLink string
			var err error

			// 获取下载链接
			if github_repository != "" {
				downloadLink, err = internal.GetGithubDownloadLink(github_repository, included_parts.Value(), excluded_parts.Value())
				if err != nil {
					log.Fatalln(err)
				}
			} else if sourceforge_rss != "" {
				downloadLink, err = internal.GetSourceForgeDownloadLink(sourceforge_rss, included_parts.Value(), excluded_parts.Value())
				if err != nil {
					log.Fatalln(err)
				}
			}

			// 下载
			if downloadLink != "" {
				err = gDownloader.Downloader(downloadLink, "", output)
				if err != nil {
					log.Fatalln(err)
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
