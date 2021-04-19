package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// 开启tcp 8000端口
	tcp, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// 接受请求
		accept, err := tcp.Accept()
		if err != nil {
			log.Printf("lost connection, %v", err)
			continue
		}

		fmt.Println("some on connect")

		// 处理每一个TCP请求
		handleConn(accept)
	}

}

// 处理客户端请求
func handleConn(con net.Conn) {
	defer con.Close()

	for {
		_, err := io.WriteString(con, time.Now().Format("15:04:05\n"))
		if err != nil {
			fmt.Println("some quit")
			return
		}
		time.Sleep(time.Millisecond * 1000)
	}
}
