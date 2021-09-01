package main

import (
	"gek_json"
	"strings"
)

type GithubAPI struct {
	TagName    string   `json:"tag_name"`
	Assets     []Assets `json:"assets"`
	TarballURL string   `json:"tarball_url"`
	ZipballURL string   `json:"zipball_url"`
	Body       string   `json:"body"`
}
type Assets struct {
	Name               string `json:"name"`
	ContentType        string `json:"content_type"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func NewGithubAPI(repo string) (githubAPI *GithubAPI, err error) {

	githubAPI = new(GithubAPI)

	// 新建json处理体
	jsoner, err := gek_json.NewJsoner(&githubAPI)
	if err != nil {
		return nil, err
	}

	// json处理体从URL中读取json数据,数据存储到githubAPI中
	err = jsoner.ReadFromURL("https://api.github.com/repos/" + repo + "/releases/latest")
	if err != nil {
		return nil, err
	}

	return githubAPI, nil
}

func (api GithubAPI) SearchRelease(part string) (downloadUrl string) {
	for _, asset := range api.Assets {
		if strings.Contains(asset.Name, part) {
			return asset.BrowserDownloadURL
		}
	}
	return ""
}
