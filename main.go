package main

import (
	"flag"
	"fmt"
	"gek_downloader"
	"gek_github"
	"log"
	"os"
)

var (
	cliRepo    string
	cliPart    string
	cliOutput  string
	cliHelp    bool
	cliVersion bool
)

func init() {
	flag.StringVar(&cliRepo, "r", "", "set repo")
	flag.StringVar(&cliPart, "p", "", "set the search part of the file name to be downloaded")
	flag.StringVar(&cliOutput, "o", "", "set output file")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `
Version:
  1.01

Usage:
  redl [Options]

Options:
  -r  <repo>    : set repo
  -p  <part>    : set the search part of the file name to be downloaded
  -o  <output>  : set output file
  -h            : show help
  -v            : show version

Example:
  1) redl -r gek64/redl -p windows-amd64
  2) redl -r gek64/redl -p windows-amd64 -o ./release-downloader-windows-amd64.exe
  3) redl -h
  4) redl -v`
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

	// 未传递repo和part参数则退出
	if cliRepo == "" || cliPart == "" {
		fmt.Println("Missing Repo or Part")
		os.Exit(0)
	}
}

func showVersion() {
	var versionInfo = `v1.01`
	fmt.Println(versionInfo)
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release
  1.01:
    - Optimized error handling`
	fmt.Println(versionInfo)
}

func main() {
	api, err := gek_github.NewGithubAPI(cliRepo)
	if err != nil {
		log.Fatalln(err)
	}

	dlUrl := api.SearchRelease(cliPart)
	if dlUrl == "" {
		log.Fatalf("Unable to find %s in release list", cliPart)
	}

	err = gek_downloader.Downloader(dlUrl, cliOutput)
	if err != nil {
		log.Fatalln(err)
	}
}
