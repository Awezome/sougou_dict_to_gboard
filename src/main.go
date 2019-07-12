package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("run scel_path")
		os.Exit(1)
	}

	root := os.Args[1]
	path := root + "/scel/"

	// define scel
	scelMap := make(map[int]string)
	scelMap[4] = "网络流行新词【官方推荐】"
	scelMap[1206] = "最新 汉语新词语选目"
	scelMap[3] = "宋词精选【官方推荐】"
	scelMap[15097] = "成语俗语【官方推荐】"
	scelMap[1] = "唐诗300首【官方推荐】"
	scelMap[15206] = "动物词汇大全【官方推荐】"
	scelMap[15128] = "法律词汇大全【官方推荐】"
	scelMap[807] = "全国省市区县地名"

	//download
	download := &Downloader{
		names: scelMap,
		path:  path,
	}
	err := download.Run()
	if err != nil {
		fmt.Println(err)
	}

	//parse
	var ws sync.WaitGroup
	dir, _ := ioutil.ReadDir(path)
	for _, fi := range dir {
		dictName := fi.Name()
		real := strings.HasSuffix(dictName, ".scel")
		if !real {
			continue
		}
		ws.Add(1)
		go func() {
			s := SougouParser{root: root}
			s.OutPutOne(path, dictName)
			defer ws.Done()
		}()
	}
	ws.Wait()
}
