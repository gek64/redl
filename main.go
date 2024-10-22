package main

import (
	"errors"
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
	var tagName string
	var included_parts cli.StringSlice
	var excluded_parts cli.StringSlice
	var no_download bool
	var output string

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "github",
			Aliases:     []string{"gh"},
			Usage:       "set github repo (example: gek64/redl)",
			Destination: &github,
		},
		&cli.StringFlag{
			Name:        "gitlab",
			Aliases:     []string{"gl"},
			Usage:       "set gitlab project id (example: 36189 or fdroid%2ffdroidclient)",
			Destination: &gitlab,
		},
		&cli.StringFlag{
			Name:        "sourceforge",
			Aliases:     []string{"sf"},
			Usage:       "set sourceforge rss url (example: https://sourceforge.net/projects/mpv-player-windows/rss?path=/64bit)",
			Destination: &sourceforge,
		},
		&cli.StringFlag{
			Name:        "tag",
			Aliases:     []string{"t"},
			Usage:       "set tag name",
			Destination: &tagName,
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
		&cli.BoolFlag{
			Name:        "no_download",
			Aliases:     []string{"nd"},
			Usage:       "output the download link without starting to download the file",
			Destination: &no_download,
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
		Version: "v2.01",
		Flags:   flags,
		Action: func(ctx *cli.Context) (err error) {
			var downloadLink string

			// 获取下载地址
			if github != "" {
				var a *internal.GithubAPI
				if tagName != "" {
					a, err = internal.GetGithubApiByTagName(github, tagName)
					if err != nil {
						return err
					}
				} else {
					a, err = internal.GetGithubApiLatest(github)
					if err != nil {
						return err
					}
				}

				downloadLink, err = a.GetDownloadLink(included_parts.Value(), excluded_parts.Value())
				if err != nil {
					return err
				}
			}
			if gitlab != "" {
				var a *internal.GitlabAPI
				if tagName != "" {
					a, err = internal.GetGitlabApiByTagName(gitlab, tagName)
					if err != nil {
						return err
					}
				} else {
					a, err = internal.GetGitlabApiLatest(gitlab)
					if err != nil {
						return err
					}
				}

				downloadLink, err = a.GetDownloadLink(included_parts.Value(), excluded_parts.Value())
				if err != nil {
					return err
				}
			}
			if sourceforge != "" {
				var a *internal.SourceForgeAPI
				if tagName != "" {
					return errors.New("sourceforge does not support searching by tag name")
				} else {
					a, err = internal.GetSourceForgeByRss(sourceforge)
					if err != nil {
						return err
					}
				}

				downloadLink, err = a.GetDownloadLink(included_parts.Value(), excluded_parts.Value())
				if err != nil {
					return err
				}
			}

			// 不进行下载文件的情况
			if no_download {
				fmt.Print(downloadLink)
				return nil
			}

			// 进行下载文件的情况
			if downloadLink != "" {
				_, err = gToolbox.CheckToolbox([]string{"curl"})
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
