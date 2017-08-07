package main

import (
	"log"
	"net"
	"strconv"
)

func runclient(port int, file string) {
	conn, err := net.Dial("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Println("无法连接,%v", err)
	}
	defer conn.Close()
	log.Println("请求成功")
}
