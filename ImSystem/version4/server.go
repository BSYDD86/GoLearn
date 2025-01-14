package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip                string
	Port              int
	OnlineUserMap     map[string]*User
	onlineUserMapLock sync.Mutex
	//消息广播
	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:            ip,
		Port:          port,
		OnlineUserMap: make(map[string]*User),
		Message:       make(chan string),
	}
	return server
}

func (this *Server) Start() {
	//socket
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

	if err != nil {
		fmt.Printf("net.Listen err:", err)
	}
	//close
	defer listener.Close()

	//启动广播
	go this.ListenMessager()
	for {
		//acccept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("listener accpet err:", err)
			continue
		}
		//dohandler
		go this.Handler(conn)
	}

}

func (this *Server) Handler(conn net.Conn) {
	//fmt.Println("current link connected")
	//用户上线
	user := NewUser(conn, this)
	user.Online()

	//接收cli发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				this.BroadCast(user, "diconnect")
				return
			}
			msgStr := string(buf[:n])
			user.Domessage(msgStr)
		}
	}()

	//阻塞，为了测试
	select {}
}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Name + "]:" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// 监听Message广播消息channel的goroutine
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		this.onlineUserMapLock.Lock()
		for _, cli := range this.OnlineUserMap {
			cli.C <- msg
		}
		this.onlineUserMapLock.Unlock()
	}
}
