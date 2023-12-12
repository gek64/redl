package internal

import (
	"fmt"
	"github.com/gek64/gek/gGithub"
	"github.com/gek64/gek/gSourceForge"
)

func GetSourceForgeDownloadLink(rssUrl string, includes []string, excludes []string) (downloadLink string, err error) {
	api, err := gSourceForge.NewAPI(rssUrl)
	if err != nil {
		return "", err
	}

	release := api.SearchRelease(includes, excludes)

	if len(release) <= 0 {
		return "", fmt.Errorf("can not find the release\n")
	}
	return release[0].Link, nil
}

func GetGithubDownloadLink(repo string, includes []string, excludes []string) (downloadLink string, err error) {
	api, err := gGithub.NewAPI(repo)
	if err != nil {
		return "", err
	}

	release := api.SearchRelease(includes, excludes)

	if len(release) <= 0 {
		return "", fmt.Errorf("can not find the release\n")
	}
	return release[0].BrowserDownloadURL, nil
}
