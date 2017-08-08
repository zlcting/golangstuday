package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func runclient(port int, file string) {
	conn, err := net.Dial("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Println("无法连接,%v", err)
	}
	defer conn.Close()
	log.Println("请求成功")

	f, err := os.Open(file)
	if err != nil {
		log.Printf("无法打开文件%s", file)
		return
	}
	defer f.Close()
	//写入头信息并等待确认
	conn.Write([]byte(file))
	p := make([]byte, 2)
	_, err = conn.Read(p)
	if err != nil {
		log.Printf("无法jieshou%v", err)
		return
	} else if string(p) != "ok" {
		log.Printf("无效的服务器响应%v", string(p))
		return
	}
	log.Println("头信息发送成功")

	_, err = io.Copy(conn, f)
	if err != nil {
		log.Printf("发送文件失败(%s):%v", file, err)
		return
	}
	log.Println("wenjian发送成功")
}
