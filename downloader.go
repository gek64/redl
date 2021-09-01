package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Downloader(url string, outputFile ...interface{}) error {
	var fileName = ""

	// 使用get方法连接url
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(response.Body)

	// outputFile 有则从中提取文件路径
	if len(outputFile) != 0 {
		fileName = outputFile[0].(string)
	}
	// outputFile为空则从header中提取文件路径
	if fileName == "" {
		contentDisposition := response.Header.Get("Content-Disposition")
		fileName = strings.Split(contentDisposition, "filename=")[1]
	}
	// header中无文件路径则使用默认路径
	if fileName == "" {
		fileName = "default_file"
	}

	// 新建输出文件
	output, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer func(output *os.File) {
		err := output.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(output)

	// 将数据写入到文件中
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}

	return nil
}
