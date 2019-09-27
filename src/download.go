package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var downPrefix = "https://pinyin.sogou.com/d/dict/download_cell.php?"

type Downloader struct {
	Id       string
	Name     string
	SavePath string
}

func (d *Downloader) One() error {
	link := downPrefix + "id=" + d.Id + "&name=" + d.Name
	name := d.SavePath + d.Name + ".scel"

	ex, err := DirExists(d.SavePath)
	if err != nil {
		exit(err)
	}
	if !ex {
		err := os.MkdirAll(d.SavePath, os.ModePerm)
		if err != nil {
			exit(err)
		}
	}

	return d.download(link, name)
}

func (d *Downloader) download(downloadUrl string, path string) error {
	fmt.Println(downloadUrl)
	res, err := http.Get(downloadUrl)
	if err != nil {
		fmt.Println("download failed")
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
