package main

import (
	"log"
	"net"
	"time"
)

/*
  创建一个tcp客户端，
  搭配服务端，用来测试tcp三次握手和四次挥手以及通信状态
  使用wireshark或者tcpdump抓包查看
*/

func main() {
	// 客户端创建一个tcp连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Printf("net dial failed.err:%+v\n", err)
		return
	}
	// 使用 conn 连接进行数据的发送和接收
	for {
		_, err = conn.Write([]byte("test socket"))
		if err != nil {
			log.Printf("send failed, err:%v\n", err)
			return
		}
		// 从服务端接收回复消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Printf("read failed:%v\n", err)
			return
		}
		log.Printf("收到服务端回复:%v\n", string(buf[:n]))

		time.Sleep(10 * time.Second)
		break
	}
	conn.Close()
}
