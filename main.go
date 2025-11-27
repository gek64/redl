package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"redl/internal"

	"github.com/unix755/xtools/xDownloader"
	"github.com/unix755/xtools/xToolbox"
	"github.com/urfave/cli/v3"
)

func main() {
	var github string
	var gitlab string
	var sourceforge string
	var tagName string
	var includedParts []string
	var excludedParts []string
	var noDownload bool
	var output string

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "github",
			Aliases:     []string{"gh"},
			Usage:       "set github repo (example: unix755/redl)",
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
			Destination: &includedParts,
		},
		&cli.StringSliceFlag{
			Name:        "excluded_parts",
			Aliases:     []string{"ep"},
			Usage:       "set release file name excluded parts",
			Destination: &excludedParts,
		},
		&cli.BoolFlag{
			Name:        "no_download",
			Aliases:     []string{"nd"},
			Usage:       "output the download link without starting to download the file",
			Destination: &noDownload,
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "set output file",
			Destination: &output,
		},
	}

	// 打印版本函数
	cli.VersionPrinter = func(cmd *cli.Command) {
		fmt.Printf("%s\n", cmd.Root().Version)
	}

	cmd := &cli.Command{
		Usage:   "Release Download Tool",
		Version: "v2.10",
		Flags:   flags,
		Action: func(ctx context.Context, cmd *cli.Command) (err error) {
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

				downloadLink, err = a.GetDownloadLink(includedParts, excludedParts)
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

				downloadLink, err = a.GetDownloadLink(includedParts, excludedParts)
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

				downloadLink, err = a.GetDownloadLink(includedParts, excludedParts)
				if err != nil {
					return err
				}
			}

			// 不进行下载文件的情况
			if noDownload {
				fmt.Print(downloadLink)
				return nil
			}

			// 进行下载文件的情况
			if downloadLink != "" {
				_, err = xToolbox.CheckToolbox([]string{"curl"})
				if err != nil {
					err = xDownloader.Download(downloadLink, output, "")
					if err != nil {
						return err
					}
				} else {
					err = xDownloader.DownloadWithCurl(downloadLink, output, "")
					if err != nil {
						return err
					}
				}
			}
			return nil
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
