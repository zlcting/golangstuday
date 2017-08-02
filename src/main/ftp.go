package main

import (
	//"flag"
	"fmt"
	"strings"
	// "io"
	// "log"
	// "net"
	// "time"
)

// func main() {

// 	listener, err := net.Listen("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("order:") //不加*号的话,输出的是内存地址

// 	var order string

// 	flag.StringVar(&order, "order", "ls", "you order")

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Print(err) // e.g., connection aborted
// 			continue
// 		}

// 		go handleConn(conn) // handle one connection at a time
// 	}
// }

func main() {
	//stack.Pop()
	//http://www.fx114.net/qa-52-177862.aspx
	var input string
	for {
		_, _ = Scanf("%s\r\n", &input)
		var eof rune = 26
		if strings.Contains(input, string(eof)) {
			Printf("输入结束\n")
			break
		}
	}
}

// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	for {
// 		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
// 		if err != nil {
// 			return // e.g., client disconnected
// 		}
// 		time.Sleep(1 * time.Second)
// 	}

// }
