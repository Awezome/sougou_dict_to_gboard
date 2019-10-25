package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var savePath string
var dictImport string
var dictTool string
var dictTxt string

func init() {
	dir, err := GetCurrentDir()
	if err != nil {
		exit("can not get current dir")
	}

	savePath = dir + "/download/"
	dictImport = dir + "/dict_with_import"
	dictTool = dir + "/dict_with_tool"
	dictTxt = dir + "/dict.txt"
}

func main() {
	dictMap := loadDictConfig()

	if err := CreateDir(savePath); err != nil {
		exit(err)
	}
	if err := CreateDir(dictImport); err != nil {
		exit(err)
	}
	if err := CreateDir(dictTool); err != nil {
		exit(err)
	}

	for id, name := range dictMap {
		worker(id, name)
	}
}

func worker(id string, name string) {
	down := &Downloader{
		Id:       id,
		Name:     name,
		SavePath: savePath,
	}
	err := down.One()
	if err != nil {
		exit(err)
	}

	dictName := savePath + name + ".scel"
	s := SougouParser{}
	err = s.OutPutOne(dictName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}

func loadDictConfig() map[string]string {
	file, err := os.Open(dictTxt)
	if err != nil {
		exit(" dictTxt 打开失败")
	}
	defer file.Close()

	br := bufio.NewReader(file)

	dictMap := make(map[string]string)
	for {
		line, _, end := br.ReadLine()
		if end == io.EOF {
			break
		}

		lineSlice := strings.Split(string(line), "|")
		if len(lineSlice) == 2 {
			dictMap[lineSlice[0]] = lineSlice[1]
		}
	}
	return dictMap
}

func exit(s interface{}) {
	fmt.Println(s)
	os.Exit(1)
}
