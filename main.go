package main

import (
	"errors"
	"fmt"
	"gboard_dict/dict"
	"io/ioutil"
	"os"

	"github.com/mholt/archiver"
	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(350, 200)
	window.SetWindowTitle("Hello Widgets Example")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Write something ...")
	widget.Layout().AddWidget(input)

	input.SetText("https://pinyin.sogou.com/dict/detail/index/4")
	widget.Layout().AddWidget(input)

	button := widgets.NewQPushButton2("and click me!", nil)
	button.ConnectClicked(func(bool) {
		err := worker(input.Text())
		if err != nil {
			widgets.QMessageBox_Information(nil, "Failed", err.Error(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		} else {
			widgets.QMessageBox_Information(nil, "OK", "成功", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		}
	})
	widget.Layout().AddWidget(button)

	window.Show()
	app.Exec()
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
	fmt.Println("finish parse ")
	return nil
}
