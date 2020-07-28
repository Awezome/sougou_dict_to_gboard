package main

import (
	"errors"
	"flag"
	"fmt"
	"gboard_dict/dict"
	"io/ioutil"
	"os"

	"github.com/mholt/archiver/v3"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return ""
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var urls arrayFlags

func init() {
	flag.Var(&urls, "url", "url")
	flag.Parse()
	if len(urls) == 0 {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	for _, url := range urls {
		err := worker(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func worker(url string) error {
	var err error

	d := &dict.Downloader{}
	url, err = d.HtmlParser(url)
	if err != nil {
		return err
	}
	if url == "" {
		return errors.New("url is empty")
	}

	bytes, err := d.GetBytes(url)
	if err != nil {
		return err
	}

	fmt.Println("start parse ")

	s := dict.SougouParser{}
	err = s.Parse(bytes)
	if err != nil {
		return err
	}

	txtPath := "./" + s.DictName + ".txt"
	zipPath := "./" + s.DictName + ".zip"
	content := s.FormatToImport()
	err = ioutil.WriteFile(txtPath, []byte(content), 0644)
	if err != nil {
		return err
	}

	os.Remove(zipPath)
	err = archiver.Archive([]string{txtPath}, zipPath)
	os.Remove(txtPath)
	if err != nil {
		return err
	}

	return nil
}
