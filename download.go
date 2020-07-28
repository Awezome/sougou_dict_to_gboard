package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Downloader struct {
	Url string
}

func (d *Downloader) GetBytes() ([]byte, error) {
	fmt.Println(d.Url)
	res, err := http.Get(d.Url)
	if err != nil {
		fmt.Println("download failed")
		return []byte{}, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
