package internal

import (
	"fmt"
	"github.com/gek64/gek/gRelease/github"
)

type GithubAPI struct {
	r *github.Release
}

func GetGithubApiLatest(repo string) (api *GithubAPI, err error) {
	r, err := github.GetReleaseLatest(repo)
	return &GithubAPI{r}, err
}

func GetGithubApiByTagName(repo string, tagName string) (api *GithubAPI, err error) {
	r, err := github.GetReleaseByTagName(repo, tagName)
	return &GithubAPI{r}, err
}

func (a *GithubAPI) GetDownloadLink(includes []string, excludes []string) (downloadLink string, err error) {
	release := a.r.GetAssets(includes, excludes)
	if len(release) <= 0 {
		return "", fmt.Errorf("can not find the release")
	}
	return release[0].BrowserDownloadURL, nil
}
