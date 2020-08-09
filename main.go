package main

import (
	"errors"
	"gboard_dict/dict"
	"io/ioutil"
	"os"

	"github.com/mholt/archiver/v3"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type CustomLabel struct {
	widgets.QLabel
	_ func()       `constructor:"init"`
	_ func(string) `signal:"updateTextFromGoroutine"`
}

func (c *CustomLabel) init() {
	c.ConnectUpdateTextFromGoroutine(c.SetText)
}

var label *CustomLabel

// type mainThreadHelper struct {
// 	core.QObject
// 	_ func(f func()) `signal:"runOnMain,auto`
// }
//func (*mainThreadHelper) runOnMain(f func()) { f() }
//var MainThreadRunner = NewMainThreadHelper(nil)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	//window.SetMinimumSize2(350, 500)
	window.SetWindowTitle("搜狗词库转Gboard词库工具")

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Write something ...")
	input.SetText("https://pinyin.sogou.com/dict/detail/index/4")

	label = NewCustomLabel(nil, 0)
	label.SetAlignment(core.Qt__AlignCenter)
	//label.SetFixedHeight(10)

	button := widgets.NewQPushButton2("Start", nil)
	button.ConnectClicked(func(bool) {
		button.SetDisabled(true)
		go func() {
			err := worker(input.Text())
			if err != nil {
				label.UpdateTextFromGoroutine(err.Error())
			} else {
				label.UpdateTextFromGoroutine("成功")
			}
			button.SetDisabled(false)
		}()
	})

	layout := widgets.NewQGridLayout2()
	layout.AddWidget3(input, 0, 0, 1, 2, 0)
	layout.AddWidget2(label, 1, 0, 0)
	layout.AddWidget2(button, 1, 1, 0)

	widget := widgets.NewQWidget(window, 0)
	widget.SetLayout(layout)
	window.SetCentralWidget(widget)
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

	label.UpdateTextFromGoroutine("start parse")

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
	label.UpdateTextFromGoroutine("finish parse")
	return nil
}
