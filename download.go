package main

import (
	"github.com/gek64/gek/gGithub"
	"github.com/gek64/gek/gSourceForge"
)

func getSourceForgeDownloadLink(rssUrl string, parts []string) (downloadLink string, err error) {
	api, err := gSourceForge.NewAPI(rssUrl)
	if err != nil {
		return "", err
	}
	return api.SearchPartsInRelease(parts)
}

func getGithubDownloadLink(repo string, parts []string) (downloadLink string, err error) {
	api, err := gGithub.NewAPI(repo)
	if err != nil {
		return "", err
	}
	return api.SearchPartsInRelease(parts)
}
