package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
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
	url, err = HtmlParser(url)
	if err != nil {
		return err
	}
	if url == "" {
		return errors.New("url is empty")
	}

	down := &Downloader{
		Url: url,
	}
	bytes, err := down.GetBytes()
	if err != nil {
		return err
	}

	s := SougouParser{}
	return s.OutPutOne(bytes)
}
