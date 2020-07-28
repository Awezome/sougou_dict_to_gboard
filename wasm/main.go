package main

import (
	"errors"
	"fmt"
	"gboard_dict/dict"
)

func main() {

}

func worker(url string) (string, error) {
	var err error

	d := &dict.Downloader{}
	url, err = d.HtmlParser(url)
	if err != nil {
		return "", err
	}
	if url == "" {
		return "", errors.New("url is empty")
	}

	bytes, err := d.GetBytes(url)
	if err != nil {
		return "", err
	}

	fmt.Println("start parse ")

	s := dict.SougouParser{}
	err = s.Parse(bytes)
	if err != nil {
		return "", err
	}

	content := s.FormatToImport()
	return content, nil
}
