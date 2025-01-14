package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
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

	for {
		//acccept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("listener accpet err:", err)
			continue
		}
		//dohandler
		go this.doHandler(conn)
	}

}

func (this *Server) doHandler(conn net.Conn) {
	fmt.Println("当前链接成功")
}
