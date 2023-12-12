package internal

import (
	"fmt"
	"github.com/gek64/gek/gGithub"
	"github.com/gek64/gek/gSourceForge"
)

func GetSourceForgeDownloadLink(releaseURL string, includes []string, excludes []string) (downloadLink string, err error) {
	r, err := gSourceForge.NewRelease(releaseURL)
	if err != nil {
		return "", err
	}

	release := r.SearchRelease(includes, excludes)

	if len(release) <= 0 {
		return "", fmt.Errorf("can not find the release")
	}
	return release[0].Link, nil
}

func GetGithubDownloadLink(repo string, includes []string, excludes []string) (downloadLink string, err error) {
	api, err := gGithub.NewReleaseLatest(repo)
	if err != nil {
		return "", err
	}

	release := api.SearchRelease(includes, excludes)

	if len(release) <= 0 {
		return "", fmt.Errorf("can not find the release")
	}
	return release[0].BrowserDownloadURL, nil
}
