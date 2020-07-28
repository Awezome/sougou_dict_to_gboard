package dict

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strings"

	"golang.org/x/text/encoding/unicode"
)

const PINGYIN_START = 0x1540
const WORD_START = 0x2628

var MAGIC = [...]byte{0x40, 0x15, 0x00, 0x00, 0x44, 0x43, 0x53, 0x01, 0x01, 0x00, 0x00, 0x00}
var PY_MAGIC = [...]byte{0x9D, 0x01, 0x00, 0x00}

type SougouParser struct {
	DictName string

	wordData  []map[string]interface{}
	pinyinMap []string
}

func (s *SougouParser) FormatToImport() string {
	content := "# Gboard Dictionary version:1\n"

	for _, line := range s.wordData {
		pinyin := strings.Join(line["p"].([]string), "")
		for _, word := range line["w"].([]string) {
			content = content + pinyin + "\t" + word + "\tzh-CN\n"
		}
	}

	//return ioutil.WriteFile(out, []byte(content), 0644)
	return content
}

func (s *SougouParser) FormatToTool() string {
	content := ""

	for _, line := range s.wordData {
		pinyin := strings.Join(line["p"].([]string), "")
		for _, word := range line["w"].([]string) {
			content = content + "[\"zh\",\"" + pinyin + "\",\"" + word + "\"],"
		}
	}
	content = "[" + strings.TrimRight(content, ",") + "]"
	return content
}

func (s *SougouParser) Parse(data []byte) error {
	flag := []byte{64, 21, 0, 0, 68, 67, 83, 1, 1, 0}
	if !bytes.Equal(flag, data[:10]) {
		return errors.New("the download file is not dict")
	}

	s.DictName = s.toString(data[0x130:0x338])
	//fmt.Println(s.dictName)

	//fmt.Println(s.toString(data[0x130:0x338]))
	//fmt.Println(toString(data[0x338:0x540]))
	//fmt.Println(toString(data[0x540:0xD40]))
	//fmt.Println(toString(data[0xD40:PINGYIN_START]))

	s.parsePinyin(data[PINGYIN_START:WORD_START])
	s.parseWord(data[WORD_START:])
	s.join()
	return nil
}

func (s *SougouParser) join() {
	pinyinMapLen := len(s.pinyinMap)
	for k, data := range s.wordData {
		newPinyin := make([]string, 0)
		pinyin := data["p"].([]int)
		for _, v := range pinyin {
			if v < pinyinMapLen {
				newPinyin = append(newPinyin, s.pinyinMap[v])
			}
		}
		s.wordData[k]["p"] = newPinyin
	}
}

func (s *SougouParser) parsePinyin(b []byte) {
	b = b[len(PY_MAGIC):]
	for len(b) > 0 {
		l := s.toInt(b[2:])
		b = b[4:]
		str := s.toString(b[:l])
		b = b[l:]
		s.pinyinMap = append(s.pinyinMap, str)
	}
}

func (s *SougouParser) parseWord(b []byte) {
	for len(b) > 0 {
		w := make([]string, 0)
		pinyin := make([]int, 0)
		pinyinWord := make(map[string]interface{})
		// 同音词
		same := s.toInt(b)
		b = b[2:]

		// 拼音
		pyLen := s.toInt(b)
		b = b[2:]
		for i := 0; i < pyLen/2; i++ {
			pinyin = append(pinyin, s.toInt(b[i*2:]))
		}
		b = b[pyLen:]

		for i := 0; i < same; i++ {
			// 词组
			wordLen := s.toInt(b)
			b = b[2:]

			if wordLen > len(b) {
				return
			}

			word := s.toString(b[:wordLen])
			b = b[wordLen:]

			// 扩展
			extLen := s.toInt(b)
			b = b[2:]
			b = b[extLen:]
			w = append(w, word)
		}
		pinyinWord["p"] = pinyin
		pinyinWord["w"] = w
		s.wordData = append(s.wordData, pinyinWord)
	}
}

func (s *SougouParser) toString(b []byte) string {
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

func (s *SougouParser) toInt(b []byte) int {
	return int(binary.LittleEndian.Uint16(b))
}
