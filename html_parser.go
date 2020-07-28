package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func HtmlParser(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	var downloadLink = ""
	doc.Find("#dict_info_dl #dict_dl_btn a").Each(func(i int, s *goquery.Selection) {
		band, ok := s.Attr("href")
		if ok {
			downloadLink = "http:" + band
		}
	})
	return downloadLink, nil
}
