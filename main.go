package main

import (
	"flag"
	"fmt"
	"github.com/gek64/gek/gDownloader"
	"log"
	"os"
)

var (
	cliRepo    string
	cliPart    string
	cliRss     string
	cliOutput  string
	cliHelp    bool
	cliVersion bool
)

func init() {
	flag.StringVar(&cliRepo, "r", "", "set repo")
	flag.StringVar(&cliPart, "p", "", "set the search part of the file name to be downloaded")
	flag.StringVar(&cliRss, "rss", "", "set rss link")
	flag.StringVar(&cliOutput, "o", "", "set output file")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
  redl {-r rope | -rss rss_link} [Options] -p [part1, part2, par3, ...]

Args:
  -r   <repo>    : set repo
  -p   <part>    : set the search part of the file name to be downloaded
  -rss <rss>     : set rss link

Options:
  -o   <output>  : set output file

Other:
  -h             : show help
  -v             : show version

Example:
  1) redl -r "gek64/redl" -p "windows-amd64"
  2) redl -r "gek64/redl" -p "windows" "amd64"
  3) redl -r "gek64/redl" -o "./release-downloader-windows-amd64.exe" -p "windows-amd64" ".exe"
  4) redl -rss "https://sourceforge.net/projects/mpv-player-windows/rss?path=/64bit" -p "x86_64" ".7z"
  5) redl -h
  6) redl -v`
		fmt.Println(helpInfo)
	}

	// 如果无 args 或者 指定 h 参数,打印用法后退出
	if len(os.Args) == 1 || cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		showVersion()
		os.Exit(0)
	}

	// 未传递repo或rss参数和part参数则退出
	if (cliRepo == "" && cliRss == "") || cliPart == "" {
		fmt.Println("Missing Repo or Rss or Part")
		os.Exit(0)
	}

	// 同时传递repo和rss参数,则退出
	if cliRepo != "" && cliRss != "" {
		fmt.Println("Rope and Rss methods can only choose one")
		os.Exit(0)
	}
}

func showVersion() {
	var versionInfo = `v1.06`
	fmt.Println(versionInfo)
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release
  1.01:
    - Optimized error handling
  1.02:
    - Add aria2, wget and build-in downloader support
  1.03:
    - Add multi-parts support
  1.04:
    - Rewrite download function
  1.05:
    - Change the description of help, parameters such as -p "windows" ".exe" -o "./bin.exe" are not supported
  1.10:
    - Add SourceForge rss support`
	fmt.Println(versionInfo)
}

func main() {
	var downloadLink string
	var err error

	// 获取下载链接
	if cliRepo != "" {
		downloadLink, err = getGithubDownloadLink(cliRepo, append(flag.Args(), cliPart))
		if err != nil {
			log.Fatalln(err)
		}
	} else if cliRss != "" {
		downloadLink, err = getSourceForgeDownloadLink(cliRss, append(flag.Args(), cliPart))
		if err != nil {
			log.Fatalln(err)
		}
	}

	// 下载
	if downloadLink != "" {
		err = gDownloader.Downloader(downloadLink, "", cliOutput)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
