# usefulTools

## pdfToVoice工具
该工具作用是阅读pdf文件。分为两步，将pdf转化为文本，然后读取文本内容
pdf文件地址支持本地pdf文件以及网络pdf文件。但网络pdf文件可能存在下载失败问题

### 使用方法：

go build main.go

main -location pdf路径 -delete true

### 代码说明：

代码分两步

pdfTransfer：负责将pdf转为txt文件。使用的是"github.com/gen2brain/go-fitz"，设置到cgo，需要设置CGO_ENABLE=1

voice：负责将pdf用语音读取出来。使用的是"github.com/go-ole/go-ole"

## socket
该工具主要是用来测试tcp连接通信的，包括三次握手、四次挥手和中间的通信状态。
包括了客户端和服务端

### 使用方法
cd socket/server

// linux

go build -o server

./server

// windows

go build -o server.exe

./server.exe

cd socket/client

// linux

go build -o client

./client

// windows

go build -o client.exe

./client.exe

### 代码说明：
就是简单的socket编程，服务端会把客户端发送的数据复制一份返回