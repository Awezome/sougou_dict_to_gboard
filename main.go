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
	window.SetFixedWidth(350)
	window.SetFixedHeight(200)
	window.SetWindowTitle("搜狗词库转Gboard词库工具")

	title := widgets.NewQLabel2("搜狗词库转Gboard词库工具V2.0", nil, 0)
	title.SetAlignment(core.Qt__AlignVCenter | core.Qt__AlignHCenter)
	title.SetStyleSheet("font-size:18px")

	info := widgets.NewQLabel2("<a href='https://github.com/Awezome/sougou_dict_to_gboard'>使用说明</a>  <a href='https://github.com/Awezome/sougou_dict_to_gboard/releases'>检查更新</a>", nil, 0)
	info.SetAlignment(core.Qt__AlignLeft)
	info.SetOpenExternalLinks(true)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("请输入搜狗词库网址...")
	input.SetText("https://pinyin.sogou.com/dict/detail/index/4")

	label = NewCustomLabel(nil, 0)
	label.SetAlignment(core.Qt__AlignLeft)
	//label.SetFixedHeight(10)

	button := widgets.NewQPushButton2("Start", nil)
	button.ConnectClicked(func(bool) {
		button.SetDisabled(true)
		go func() {
			err := worker(input.Text())
			if err != nil {
				label.UpdateTextFromGoroutine(err.Error())
			}
			button.SetDisabled(false)
		}()
	})

	layout := widgets.NewQGridLayout2()
	layout.AddWidget3(title, 0, 0, 3, 2, 0)
	layout.AddWidget3(input, 1, 0, 1, 2, 0)
	layout.AddWidget2(label, 2, 0, 0)
	layout.AddWidget2(button, 3, 1, 0)
	layout.AddWidget2(info, 4, 0, 0)

	widget := widgets.NewQWidget(window, 0)
	widget.SetLayout(layout)
	window.SetCentralWidget(widget)
	window.Show()
	app.Exec()
}

func worker(url string) error {
	var err error
	label.UpdateTextFromGoroutine("开始加载...")
	d := &dict.Downloader{}
	url, err = d.HtmlParser(url)
	if err != nil {
		return err
	}
	if url == "" {
		return errors.New("url is empty")
	}
	label.UpdateTextFromGoroutine("解析词库下载链接...")

	bytes, err := d.GetBytes(url)
	if err != nil {
		return err
	}

	label.UpdateTextFromGoroutine("转换中...")

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
	label.UpdateTextFromGoroutine("转换完成")
	return nil
}
