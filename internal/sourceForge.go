package internal

import (
	"fmt"

	"github.com/unix755/xtools/xRelease/sourceForge"
)

type SourceForgeAPI struct {
	r *sourceForge.Release
}

func GetSourceForgeByRss(rssUrl string) (api *SourceForgeAPI, err error) {
	r, err := sourceForge.GetReleaseByRss(rssUrl)
	return &SourceForgeAPI{r}, err
}

func (a *SourceForgeAPI) GetDownloadLink(includes []string, excludes []string) (downloadLink string, err error) {
	release := a.r.GetAssets(includes, excludes)
	if len(release) <= 0 {
		return "", fmt.Errorf("can not find the release")
	}
	return release[0].Link, nil
}
