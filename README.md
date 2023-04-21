# usefulTools

pdfToVoice工具
该工具作用是阅读pdf文件。分为两步，将pdf转化为文本，然后读取文本内容
pdf文件地址支持本地pdf文件以及网络pdf文件。但网络pdf文件可能存在下载失败问题

使用方法：

go build main.go

main -location pdf路径 -delete true

代码说明：

代码分两步

pdfTransfer：负责将pdf转为txt文件。使用的是"github.com/gen2brain/go-fitz"，设置到cgo，需要设置CGO_ENABLE=1

voice：负责将pdf用语音读取出来。使用的是"github.com/go-ole/go-ole"
