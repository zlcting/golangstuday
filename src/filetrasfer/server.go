package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func handle(conn net.Conn) {
	defer conn.Close()
	remoteAdrr := conn.RemoteAddr().String()
	log.Println("ip address:", remoteAdrr)
	p := make([]byte, 1024)
	n, err := conn.Read(p)
	if err != nil {
		log.Printf("读取文件头失败(%s)%v", remoteAdrr, err)
		return
	} else if n == 0 {
		log.Printf("空文件头(%s)", remoteAdrr)
		return
	}
	fileName := string(p[:n])
	log.Printf("文件(%s)", fileName)
	conn.Write([]byte("ok"))

	//打开一个本地文件
	os.MkdirAll("receive", os.ModePerm)
	f, err := os.Create("receive/" + fileName)
	if err != nil {
		log.Printf("无法创建文件(%s):%v", remoteAdrr, fileName)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, conn)
	for {
		buffer := make([]byte, 1024*200) //每次读取200字节可自定义
		_, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			log.Printf("读取失败(%s):%v", remoteAdrr, err)
			return
		} else if err == io.EOF {
			break
		}
	}
	if err != nil {
		log.Printf("文件接收失败(%s):%v", remoteAdrr, err)
	}
	log.Printf("文件接收成功(%s):%s", remoteAdrr, fileName)
}

func runserver(port int) {
	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("监听失败:%v", err)
	}
	log.Println("服务器已启动")
	for {
		conn, err := l.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); !ok || !ne.Temporary() {
				log.Println("请求失败", err)
			}
			continue

		}
		log.Println("请求接收成功")
		handle(conn)
	}
}
