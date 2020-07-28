package dict

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Downloader struct{}

func (d *Downloader) GetBytes(url string) ([]byte, error) {
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("download failed")
		return []byte{}, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (d *Downloader) HtmlParser(url string) (string, error) {
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
