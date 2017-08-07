package main

import (
	"log"
	"net"
	"strconv"
)

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
		_ = conn
	}
}
