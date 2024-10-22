package internal

import (
	"fmt"
	"github.com/gek64/gek/gRelease/gitlab"
)

type GitlabAPI struct {
	r *gitlab.Release
}

func GetGitlabApiLatest(repo string) (api *GitlabAPI, err error) {
	r, err := gitlab.GetReleaseLatest(repo)
	return &GitlabAPI{r}, err
}

func GetGitlabApiByTagName(repo string, tagName string) (api *GitlabAPI, err error) {
	r, err := gitlab.GetReleaseByTagName(repo, tagName)
	return &GitlabAPI{r}, err
}

func (a *GitlabAPI) GetDownloadLink(includes []string, excludes []string) (downloadLink string, err error) {
	release := a.r.GetAssets(includes, excludes)
	if len(release) <= 0 {
		return "", fmt.Errorf("can not find the release")
	}
	return release[0].URL, nil
}
