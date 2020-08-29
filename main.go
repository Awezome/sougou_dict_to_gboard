package main

import (
	"errors"
	"fmt"
	"gboard_dict/dict"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/mholt/archiver/v3"
)

var LabelInfo = widget.NewLabel("")

var ButtonStart = &widget.Button{}

func main() {
	app := app.New()
	ShowUI(app)
	app.Run()
}

func ShowUI(app fyne.App) {
	window := app.NewWindow("搜狗词库转Gboard工具")

	link, _ := url.Parse("https://github.com/Awezome/sougou_dict_to_gboard")

	input := widget.NewEntry()
	input.SetText("https://pinyin.sogou.com/dict/detail/index/4")

	ButtonStart.Text = "Start"
	ButtonStart.OnTapped = func() {
		go func() {
			ButtonStart.Disable()
			err := worker(input.Text)
			if err != nil {
				writeMessage(err.Error())
			}
			ButtonStart.Enable()
		}()
	}
	ButtonStart.ExtendBaseWidget(ButtonStart)

	window.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		widget.NewLabel("Sougou Dict to Gboard"),
		input,
		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			LabelInfo,
			widget.NewLabel(""),
			widget.NewLabel(""),
			ButtonStart,
		),
		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			widget.NewHyperlink("version v2.1", link),
			widget.NewLabel(""),
			widget.NewLabel(""),
			widget.NewLabel(""),
		),
	),
	)
	window.Resize(fyne.NewSize(440, 160))
	window.SetFixedSize(true)
	window.Show()
}

func worker(url string) error {
	var err error
	writeMessage("load...")
	fmt.Println(1)
	d := &dict.Downloader{}
	fmt.Println(2)
	url, err = d.HtmlParser(url)
	if err != nil {
		return err
	}
	if url == "" {
		return errors.New("url is empty")
	}
	writeMessage("download...")
	bytes, err := d.GetBytes(url)
	if err != nil {
		return err
	}

	writeMessage("read...")

	s := dict.SougouParser{}
	err = s.Parse(bytes)
	if err != nil {
		return err
	}
	writeMessage("parse...")

	dir, _ := os.Getwd()
	txtPath := filepath.Join(dir, s.DictName+".txt")
	zipPath := filepath.Join(dir, s.DictName+".zip")
	content := s.FormatToImport()
	err = ioutil.WriteFile(txtPath, []byte(content), 0644)
	if err != nil {
		return err
	}
	writeMessage("zip...")
	os.Remove(zipPath)
	err = archiver.Archive([]string{txtPath}, zipPath)
	os.Remove(txtPath)
	if err != nil {
		return err
	}
	writeMessage("finish")
	return nil
}

func writeMessage(m string) {
	LabelInfo.SetText(m)
}
