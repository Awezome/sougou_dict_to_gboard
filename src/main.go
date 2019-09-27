package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dictMap := loadDictConfig()

	savePath := "./download/"
	for id, name := range dictMap {
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
		s := SougouParser{root: "."}
		s.OutPutOne(dictName)
	}
}

func loadDictConfig() map[string]string {
	file, err := os.Open("./dict.txt")
	if err != nil {
		exit("文件打开失败")
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
