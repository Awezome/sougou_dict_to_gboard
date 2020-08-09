package main

import (
	"errors"
	"fmt"
	"gboard_dict/dict"
	"io/ioutil"
	"log"
	"os"

	"github.com/mholt/archiver/v3"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

var dom *sciter.Element
var domMessage *sciter.Element

func main() {
	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG, nil)
	if err != nil {
		log.Fatal("Create Window Error: ", err)
	}
	w.LoadFile("index.html")

	dom, _ = w.GetRootElement()
	domMessage, _ = dom.SelectFirst("#message")

	setEventHandler(w)
	w.Show()
	w.Run()
}

func setEventHandler(w *window.Window) {

	w.DefineFunction("getNetInformation", func(args ...*sciter.Value) *sciter.Value {
		url := args[0].String()
		fmt.Println(url)

		go func() {
			err := worker(url)
			if err != nil {
				writeMessage(err.Error())
			}
		}()

		return sciter.NullValue()
	})
}

func worker(url string) error {
	var err error
	writeMessage("开始加载...")
	d := &dict.Downloader{}
	url, err = d.HtmlParser(url)
	if err != nil {
		return err
	}
	if url == "" {
		return errors.New("url is empty")
	}
	writeMessage("下载...")
	bytes, err := d.GetBytes(url)
	if err != nil {
		return err
	}

	writeMessage("读取...")

	s := dict.SougouParser{}
	err = s.Parse(bytes)
	if err != nil {
		return err
	}
	writeMessage("生成词库文本...")
	txtPath := "./" + s.DictName + ".txt"
	zipPath := "./" + s.DictName + ".zip"
	content := s.FormatToImport()
	err = ioutil.WriteFile(txtPath, []byte(content), 0644)
	if err != nil {
		return err
	}
	writeMessage("开始打包...")
	os.Remove(zipPath)
	err = archiver.Archive([]string{txtPath}, zipPath)
	os.Remove(txtPath)
	if err != nil {
		return err
	}
	writeMessage("完成")
	return nil
}

func writeMessage(m string) {
	domMessage.SetHtml(m, sciter.SIH_REPLACE_CONTENT)
}
