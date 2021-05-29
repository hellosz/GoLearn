package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// 	启动/监听 tcp 服务
	conn, err := net.Listen("tcp", "127.0.0.1:8010")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go broadcaster()

	// 处理TCP连接
	for {
		con, err := conn.Accept()
		fmt.Println("接收到连接请求")
		if err != nil {
			log.Println("处理网络连接失败, 原因：" + err.Error())
			continue
		}

		go handleconn(con)
	}
}

// 用来接收消息的队列(指定只接收消息)
type client chan<- string

// 批量定义变量
var (
	arriving = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

// broadcaster 进行任务调度
func broadcaster() {
	// 定义客户端通道池子
	clients := make(map[client]bool)

	// 循环进行任务调度
	for {
		select {
		case msg := <-message:
			fmt.Printf("接收到消息:%s, 进行广播\n", msg)
			for cli := range clients {
				cli <- msg + "\n"
			}
		case cli := <-arriving:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}

	}
}

// 处理数据库链接
func handleconn(con net.Conn) {
	// 注册，将通道进行拆分（只接收和只发送消息）
	ch := make(chan string)
	go conwriter(con, ch)
	who := con.RemoteAddr().String()
	message <- who + " has arrived\n"
	ch <- fmt.Sprintf("you are %s\n", who)
	arriving <- ch

	// 接收消息，触发广播
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		message <- scanner.Text()
	}

	// 注销
	message <- who + " has left\n"
	leaving <- ch
	con.Close()
}

// conwriter 将数据输出到客户端(只接收chan 接收消息)
func conwriter(con net.Conn, ch <-chan string) {
	for mes := range ch {
		_, err := fmt.Fprintf(con, mes)
		if err != nil {
			log.Println("写消息给客户端失败，原因：" + mes)
		}
	}
}
