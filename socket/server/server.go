package main

import (
	"bufio"
	"log"
	"net"
)

/*
  创建一个tcp服务端，
  搭配客户端，用来测试tcp三次握手和四次挥手以及通信状态
  使用wireshark或者tcpdump抓包查看
*/

func process(conn net.Conn) {
	// 处理完关闭连接
	defer conn.Close()

	// 针对当前连接做发送和接受操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			log.Printf("read from conn failed, err:%v\n", err)
			break
		}

		recv := string(buf[:n])
		log.Printf("收到的数据：%v\n", recv)

		// 将接受到的数据返回给客户端
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			log.Printf("write from conn failed, err:%v\n", err)
			break
		}
	}
}

func main() {
	// 服务端监听8080端口号
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("net listen failed.err:%+v", err)
		return
	}
	for {
		// 接收来自客户端的连接
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis accept failed.err:%+v", err)
			break
		}
		go process(conn)
	}
}
