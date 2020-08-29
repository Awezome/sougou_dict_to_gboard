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

const html = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <style>table {width: 100%;border-collapse: collapse;}table td {padding: 2px 0;height: 32px;border: 0 solid #ccc;}</style>
</head>

<body>
    <table>
        <tr>
            <td colspan="2" style="text-align: center;height: 36px;">
                <span style="font-size: 16px;">搜狗词库转Gboard工具</span>
            </td>
        </tr>
        <tr>
            <td colspan="2">
                <input style="width: 100%;" type="text" name="url"
                    value="https://pinyin.sogou.com/dict/detail/index/4" />
            </td>
        </tr>
        <tr>
            <td>
                <div id="message"></div>
            </td>
            <td style="text-align: right;">
                <button id="btn">Start</button>
            </td>
        </tr>
        <tr>
            <td colspan="2" style="font-size: 12px;">
                当前版本v2.0
                <a href='https://github.com/Awezome/sougou_dict_to_gboard/releases'>检查更新</a>
                <a href='https://github.com/Awezome/sougou_dict_to_gboard'>使用说明</a>
            </td>
        </tr>
    </table>
</body>
<script type="text/tiscript">
$(#btn).on("click",function(){
    var url=$(input[name="url"]).value.trim()
    view.getNetInformation(url);
});
self.on("click", "a[href^=http]", function(evt) {
    var href = this.attributes["href"];
    Sciter.launch(href);
    return true;
  });
</script>
</html>
`

func main() {
	win, err := window.New(sciter.SW_TITLEBAR|sciter.SW_MAIN|sciter.SW_CONTROLS,
		&sciter.Rect{Left: 100, Top: 100, Right: 450, Bottom: 300})
	if err != nil {
		log.Fatal("Create Window Error: ", err)
	}
	win.SetOption(sciter.SCITER_SET_SCRIPT_RUNTIME_FEATURES, sciter.ALLOW_SYSINFO)
	win.LoadHtml(html, "")

	dom, _ = win.GetRootElement()
	domMessage, _ = dom.SelectFirst("#message")

	setEventHandler(win)
	win.AddQuitMenu()
	win.Show()
	win.Run()
}

func setEventHandler(w *window.Window) {
	w.DefineFunction("getNetInformation", func(args ...*sciter.Value) *sciter.Value {
		url := args[0].String()
		fmt.Println(url)

		domButton, _ := dom.SelectFirst("#btn")
		domButton.SetState(sciter.STATE_DISABLED, 0, true)
		go func() {
			err := worker(url)
			if err != nil {
				writeMessage(err.Error())
			}
			domButton.SetState(0, sciter.STATE_DISABLED, true)
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
	txtPath := s.DictName + ".txt"
	zipPath := s.DictName + ".zip"
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
