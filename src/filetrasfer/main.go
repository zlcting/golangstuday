package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	model = flag.String("model", "", "运行模式")
	port  = flag.Int("port", 5050, "服务器监听端口")
	file  = flag.String("file", "", "文件名称")
)

func main() {
	flag.Parse()
	fmt.Println(*model, *port, *file)
	//check model
	switch *model {
	case "server":
		runserver(*port)
	case "client":
		runclient(*port, *file)
	default:
		log.Fatalf("未知运行模式,%v", *model)
	}
}
