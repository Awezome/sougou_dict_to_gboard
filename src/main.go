package main

import (
	"archive/zip"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/unicode"
)

const PINGYIN_START = 0x1540
const WORD_START = 0x2628

var MAGIC = [...]byte{0x40, 0x15, 0x00, 0x00, 0x44, 0x43, 0x53, 0x01, 0x01, 0x00, 0x00, 0x00}
var PY_MAGIC = [...]byte{0x9D, 0x01, 0x00, 0x00}

type PinyinWord []map[string]interface{}

func main() {
	root := os.Args[1]
	path := root + "/scel/"
	dir, _ := ioutil.ReadDir(path)
	for _, fi := range dir {
		real := strings.HasSuffix(fi.Name(), ".scel")
		if !real {
			continue
		}
		fileName := path + fi.Name()
		content, _ := ioutil.ReadFile(fileName)
		wordData := parse(content)

		dictPath := root + "/dict_with_tool/" + fi.Name() + ".txt"
		outputToGboardTool(wordData, dictPath)

		dictPath = root + "/dict_with_import/dictionary.txt"
		zipPath := root + "/dict_with_import/" + fi.Name() + ".zip"
		outputToGboardImport(wordData, dictPath)
		os.Remove(zipPath)
		zipFile(dictPath, zipPath)
		os.Remove(dictPath)
	}
}

func zipFile(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+"/")
		// header.Name = path
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}

func outputToGboardImport(wordData PinyinWord, out string) {
	content := "# Gboard Dictionary version:1\n"

	for _, line := range wordData {
		pinyin := strings.Join(line["p"].([]string), "")
		for _, word := range line["w"].([]string) {
			content = content + pinyin + "\t" + word + "\tzh-CN\n"
		}
	}

	err := ioutil.WriteFile(out, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func outputToGboardTool(wordData PinyinWord, out string) {
	content := ""

	for _, line := range wordData {
		pinyin := strings.Join(line["p"].([]string), "")
		for _, word := range line["w"].([]string) {
			content = content + "[\"zh\",\"" + pinyin + "\",\"" + word + "\"],"
		}
	}
	content = "[" + strings.TrimRight(content, ",") + "]"
	err := ioutil.WriteFile(out, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func parse(data []byte) PinyinWord {
	fmt.Println(toString(data[0x130:0x338]))
	//fmt.Println(toString(data[0x338:0x540]))
	//fmt.Println(toString(data[0x540:0xD40]))
	//fmt.Println(toString(data[0xD40:PINGYIN_START]))

	pinyinMap := make([]string, 0)
	wordData := make(PinyinWord, 0)
	parsePinyin(data[PINGYIN_START:WORD_START], &pinyinMap)
	parseWord(data[WORD_START:], &wordData)
	join(pinyinMap, wordData)
	return wordData
}

func join(pinyinMap []string, wordData PinyinWord) {
	for k, data := range wordData {
		newPinyin := make([]string, 0)
		pinyin := data["p"].([]int)
		for _, v := range pinyin {
			newPinyin = append(newPinyin, pinyinMap[v])
		}
		wordData[k]["p"] = newPinyin
	}
}

func parsePinyin(b []byte, pinyinMap *[]string) {
	b = b[len(PY_MAGIC):]
	for len(b) > 0 {
		l := toInt(b[2:])
		b = b[4:]
		s := toString(b[:l])
		b = b[l:]
		*pinyinMap = append(*pinyinMap, s)
	}
}

func parseWord(b []byte, wordData *PinyinWord) {
	for len(b) > 0 {
		w := make([]string, 0)
		pinyin := make([]int, 0)
		pinyinWord := make(map[string]interface{})
		// 同音词
		same := toInt(b)
		b = b[2:]

		// 拼音
		pyLen := toInt(b)
		b = b[2:]
		for i := 0; i < pyLen/2; i++ {
			pinyin = append(pinyin, toInt(b[i*2:]))
		}
		b = b[pyLen:]

		for i := 0; i < same; i++ {
			// 词组
			wordLen := toInt(b)
			b = b[2:]

			if wordLen > len(b) {
				return
			}

			word := toString(b[:wordLen])
			b = b[wordLen:]

			// 扩展
			extLen := toInt(b)
			b = b[2:]
			b = b[extLen:]
			w = append(w, word)
		}
		pinyinWord["p"] = pinyin
		pinyinWord["w"] = w
		*wordData = append(*wordData, pinyinWord)
	}
}

func toString(b []byte) string {
	i := 0
	for ; i < len(b); i += 2 {
		if b[i] == 0 && b[i+1] == 0 {
			break
		}
	}

	decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	utf8, _ := decoder.Bytes(b[:i])
	return string(utf8)
}

func toInt(b []byte) int {
	return int(binary.LittleEndian.Uint16(b))
}
