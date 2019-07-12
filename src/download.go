package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Downloader struct {
	names map[int]string
	path  string
}

func (d *Downloader) Run() error {
	var wg sync.WaitGroup
	for id, name := range d.names {
		urlName := url.QueryEscape(strings.Trim(name, " "))
		downloadUrl := "https://pinyin.sogou.com/d/dict/download_cell.php?id=" + strconv.Itoa(id) + "&name=" + urlName + "&f=detail"
		path := d.path + name + ".scel"

		wg.Add(1)
		go func() error {
			err := d.download(downloadUrl, path)
			defer wg.Done()
			return err
		}()
	}
	wg.Wait()
	return nil
}

func (d *Downloader) download(downloadUrl string, path string) error {
	res, err := http.Get(downloadUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return err
	}
	defer res.Body.Close()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, res.Body)
	fmt.Println("download finish " + path)
	return err
}
