# 搜狗词库转手机Gboard词库工具

本工具可以自动下载并转换搜狗词库到Gboard词库

# 使用说明
1. 在搜狗词库找到你想要转的词库 https://pinyin.sogou.com/dict/
2. 比如，网络流行新词，它的网址是 https://pinyin.sogou.com/dict/detail/index/4  ，那么下载id就是4
3. 在程序所在目录的 dict.txt 里写上4英文竖线|和词库的名字，注意不要有空格 比如
```
4|网络流行新词
```
4. 一行一个，多个词库就多行，比如
```
4|网络流行新词
1206|最新汉语新词语选目
3|宋词精选
15097|成语俗语
1|唐诗300首
15206|动物词汇大全
15128|法律词汇大全
807|全国省市区县地名
470|重庆方言
265|重庆区域地名
```
5. windows 双击运行 sougou_to_gboard.ext , mac 用命令行运行 sougou_to_gboard
6. 生成了来两个目录，一个是dict_with_import，里面的词库可以用gboard自带的导入功能导入。另一个是dict_with_tool，里面用工具导入，方法见酷安gboard下的热贴。
